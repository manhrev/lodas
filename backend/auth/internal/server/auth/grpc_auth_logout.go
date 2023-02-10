package auth

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authServer) LogOut(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {

	if err := s.cache.AccessToken().Revoke(ctx, s.extractor.GetTokenID(ctx)); err != nil {
		log.Fatalf("could not set cache token: %v", err)
	}

	return &emptypb.Empty{}, nil
}
