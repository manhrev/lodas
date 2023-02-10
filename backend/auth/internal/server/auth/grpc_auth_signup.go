package auth

import (
	"context"

	auth "github.com/manhrev/runtracking/backend/auth/pkg/api"
)

func (s *authServer) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.SignUpReply, error) {
	reply, tokens, err := s.signUp.SignUp(ctx, request)
	if err != nil {
		return nil, err
	}

	s.cacheTokens(ctx, tokens)

	return reply, nil
}
