package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/manhrev/runtracking/backend/auth/config"
	"github.com/manhrev/runtracking/backend/auth/internal/cache"
	"github.com/manhrev/runtracking/backend/auth/internal/feature/signin"
	"github.com/manhrev/runtracking/backend/auth/internal/feature/signup"
	"github.com/manhrev/runtracking/backend/auth/internal/server/auth"
	authz "github.com/manhrev/runtracking/backend/auth/internal/server/authz"
	"github.com/manhrev/runtracking/backend/auth/internal/service/token"
	pb "github.com/manhrev/runtracking/backend/auth/pkg/api"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
	"github.com/manhrev/runtracking/backend/auth/pkg/extractor"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	ENVIRONMENT_MODE = "ENVIRONMENT_MODE"
)

func Run() {
	server := newServer()
	Serve(server)
}

func newServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}

func Serve(server *grpc.Server) {
	conf := config.GetConfig(os.Getenv(ENVIRONMENT_MODE))
	configuration, err := conf.New()
	if err != nil {
		log.Fatalf("cannot read values from yml config file: %v", err)
	}

	// init other services client connections, database driver and pass to server
	db := configuration.Database
	entClient, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", db.Username, db.Password, db.Domain, db.Port, db.Name))
	if err != nil {
		log.Fatalf("cannot create database connection: %v", err)
	}
	defer func() {
		if err := entClient.Close(); err != nil {
			panic(err)
		}
	}()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	tokenService, err := token.New()
	if err != nil {
		log.Fatalf("cannot create token service: %v", err)
	}

	extractorService := extractor.New()

	featureSignIn, err := signin.New(entClient, tokenService)
	if err != nil {
		log.Fatalf("cannot create featureSignIn: %v", err)
	}

	featureSignUp, err := signup.New(entClient, tokenService)
	if err != nil {
		log.Fatalf("cannot create featureSignUp: %v", err)
	}

	cacheService, err := cache.New(entClient, configuration)

	if err != nil {
		log.Fatalf("cannot create cacheService: %v", err)
	}

	logger, _ := zap.NewProduction()

	// register main and other server servers
	authv3.RegisterAuthorizationServer(server, authz.NewServer(tokenService, cacheService))

	if err != nil {
		logger.Fatal("can not create cache", zap.Error(err))
	}

	// Register Auth Server
	pb.RegisterAuthServer(server, auth.NewServer(entClient, tokenService, featureSignIn, featureSignUp, cacheService, extractorService))
	// Register User Server

	grpc := configuration.Grpc
	lis, err := net.Listen(grpc.Network, fmt.Sprintf("%s:%s", grpc.Host, grpc.Port))
	if err != nil {
		log.Fatalf("error while create listen: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
