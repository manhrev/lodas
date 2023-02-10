package auth

import (
	"context"

	auth "github.com/manhrev/lodas/backend/auth/pkg/api"
)

func (s *authServer) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginReply, error) {
	reply, tokens, err := s.signIn.SignIn(ctx, request)
	if err != nil {
		return reply, err
	}

	s.cacheTokens(ctx, tokens)

	return &auth.LoginReply{
		UserId:      reply.UserId,
		AccessToken: tokens.AccessToken.Raw,
	}, nil
}
