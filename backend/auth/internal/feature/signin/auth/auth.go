package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/manhrev/runtracking/backend/auth/internal/service/token"
	pb "github.com/manhrev/runtracking/backend/auth/pkg/api"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent"
	"github.com/manhrev/runtracking/backend/auth/pkg/ent/user"
)

type Auth interface {
	SignIn(ctx context.Context, request *pb.LoginRequest) (*pb.LoginReply, *token.Tokens, error)
}

type auth struct {
	ent   *ent.Client
	token token.Token
}

func New(
	ent *ent.Client,
	token token.Token) (Auth, error) {

	return &auth{
		ent:   ent,
		token: token,
	}, nil
}

func (s auth) SignIn(ctx context.Context, request *pb.LoginRequest) (*pb.LoginReply, *token.Tokens, error) {

	username := strings.ToLower(request.GetUserName())
	user, err := s.ent.User.Query().
		Where(user.Username(username)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil, err
		}
		return nil, nil, err
	}

	if user.Password != request.GetPassword() {
		log.Println("User password don't match")
		return nil, nil, errors.New("User password don't match")
	}

	tokens, err := s.token.Create(user.ID)

	if err != nil {
		log.Println("Can't create access token")
		return nil, nil, errors.New("Can't create access token")
	}

	// fmt.Printf("Tokens: %s", tokens.AccessToken.Raw)

	return &pb.LoginReply{
		UserId:      user.ID,
		AccessToken: tokens.AccessToken.Raw,
	}, tokens, nil
}

func errorWithDefaultMessage(err error) error {
	// errCode := status.Code(err)
	// if errCode == code.Code_MY_ID_SIGN_IN_CREDENTIALS_IS_INVALID {
	// 	return status.New(errCode, status.DefaultInvalidMessage)
	// }

	return err
}
