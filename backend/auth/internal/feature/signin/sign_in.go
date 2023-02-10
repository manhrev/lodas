package signin

import (
	"context"

	"github.com/manhrev/runtracking/backend/auth/internal/feature/signin/auth"
	"github.com/manhrev/runtracking/backend/auth/internal/service/token"
	pb "github.com/manhrev/runtracking/backend/auth/pkg/api"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
)

type SignIn interface {
	SignIn(ctx context.Context, request *pb.LoginRequest) (*pb.LoginReply, *token.Tokens, error)
}

type signIn struct {
	auth auth.Auth
}

func New(
	ent *ent.Client,
	token token.Token) (SignIn, error) {
	auth, err := auth.New(ent, token)
	if err != nil {
		return nil, err
	}

	return &signIn{
		auth: auth,
	}, nil
}

func (s signIn) SignIn(ctx context.Context, request *pb.LoginRequest) (*pb.LoginReply, *token.Tokens, error) {
	return s.auth.SignIn(ctx, request)
}
