package token

import (
	"github.com/manhrev/lodas/backend/auth/internal/service/token/signer"
)

type Tokens struct {
	UserID      int64
	AccessToken *signer.Token
}

type Token interface {
	AccessToken() signer.Signer
	Create(userID int64) (tokens *Tokens, err error)
}

type token struct {
	accessToken signer.Signer
}

func New() (Token, error) {
	accessToken, err := signer.New()
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &token{
		accessToken: accessToken,
	}, nil
}

func (t *token) AccessToken() signer.Signer {
	return t.accessToken
}

// Create TODO audience
func (t *token) Create(userID int64) (*Tokens, error) {
	accessToken, err := t.accessToken.Create(userID)
	if err != nil {
		return nil, err
	}

	return &Tokens{userID, accessToken}, nil
}
