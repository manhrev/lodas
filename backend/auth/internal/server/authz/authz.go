package authz

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	authv3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/manhrev/lodas/backend/auth/internal/cache"
	"github.com/manhrev/lodas/backend/auth/internal/service/token"
	"github.com/manhrev/lodas/backend/auth/internal/service/token/signer"
	"github.com/manhrev/lodas/backend/auth/pkg/data"
	"github.com/manhrev/lodas/backend/auth/pkg/header"

	typev3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	_bearer        = "Bearer"
	_authorization = "authorization"
)

type authZServer struct {
	token token.Token
	cache *cache.Cache
}

func NewServer(token token.Token, cache *cache.Cache) authv3.AuthorizationServer {
	return &authZServer{
		token: token,
		cache: cache,
	}
}

func (s *authZServer) Check(ctx context.Context, req *authv3.CheckRequest) (*authv3.CheckResponse, error) {
	authorization := req.Attributes.Request.Http.Headers[_authorization]
	extracted := strings.Fields(authorization)
	if len(extracted) == 2 && extracted[0] == _bearer {
		tk, u, err := s.verifyAccessToken(ctx, extracted[1])
		if err == nil {
			return &authv3.CheckResponse{
				HttpResponse: &authv3.CheckResponse_OkResponse{
					OkResponse: &authv3.OkHttpResponse{
						Headers: s.createHeaders(tk, u),
					},
				},
				Status: &status.Status{
					Code: int32(rpc.OK),
				},
			}, nil
		} else {
			log.Println("verify token error")
			if err == errTokenInactive {
				// Using StatusCode_Gone for the reason that kick by another device
				return buildDeniedResponse(int32(rpc.UNAUTHENTICATED), typev3.StatusCode_Conflict), nil
			}
		}
	} else {
		log.Println("invalid header")
	}

	return buildDeniedResponse(int32(rpc.UNAUTHENTICATED), typev3.StatusCode_Unauthorized), nil
}

func buildDeniedResponse(outerCode int32, innerCode typev3.StatusCode) *authv3.CheckResponse {
	return &authv3.CheckResponse{
		Status: &status.Status{
			Code: outerCode,
		},
		HttpResponse: &authv3.CheckResponse_DeniedResponse{
			DeniedResponse: &authv3.DeniedHttpResponse{
				Status: &typev3.HttpStatus{
					Code: innerCode,
				},
			},
		},
	}
}

var (
	errCouldNotParseToken = errors.New("could not parse token")
	errTokenInvalid       = errors.New("invalid token")
	errCacheNotFound      = errors.New("cache token not found")
	errTokenInactive      = errors.New("inactive token")
	errUserNotFound       = errors.New("user not found")
)

func (s *authZServer) verifyAccessToken(ctx context.Context, accessToken string) (*signer.Token, *data.User, error) {
	tk, err := s.token.AccessToken().Parse(accessToken)
	if err != nil {
		return nil, nil, errCouldNotParseToken
	}
	if tk.Valid() != nil {
		return nil, nil, errTokenInvalid
	}

	active, err := s.cache.AccessToken().IsActive(ctx, tk)
	if err != nil {
		return nil, nil, errCacheNotFound
	}
	if !active {
		return nil, nil, errTokenInactive
	}

	u, err := s.cache.User().Get(ctx, tk.UserID)
	if err != nil {
		return nil, nil, errUserNotFound
	}

	return tk, u, nil
}

func (s *authZServer) createHeaders(tk *signer.Token, user *data.User) []*corev3.HeaderValueOption {
	headers := []*corev3.HeaderValueOption{
		{
			Append: &wrapperspb.BoolValue{Value: false},
			Header: &corev3.HeaderValue{
				Key:   header.TokenID,
				Value: tk.Id,
			},
		},
		{
			Append: &wrapperspb.BoolValue{Value: false},
			Header: &corev3.HeaderValue{
				Key:   header.UserID,
				Value: fmt.Sprint(user.GetUserId()),
			},
		},
	}

	return headers
}
