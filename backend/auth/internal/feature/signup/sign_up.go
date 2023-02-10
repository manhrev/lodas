package signup

import (
	"context"

	"github.com/manhrev/runtracking/backend/auth/internal/feature/signup/auth"
	"github.com/manhrev/runtracking/backend/auth/internal/service/token"
	pb "github.com/manhrev/runtracking/backend/auth/pkg/api"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
)

type SignUp interface {
	SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpReply, *token.Tokens, error)
}

type signUp struct {
	auth auth.Auth
}

func New(
	ent *ent.Client,
	token token.Token) (SignUp, error) {
	auth, err := auth.New(ent, token)
	if err != nil {
		return nil, err
	}

	return &signUp{
		auth: auth,
	}, nil
}

func (s signUp) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpReply, *token.Tokens, error) {
	return s.auth.SignUp(ctx, request)
}
