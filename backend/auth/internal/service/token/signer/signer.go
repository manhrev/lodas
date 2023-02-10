package signer

import (
	"crypto/ed25519"
	"encoding/pem"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

type Signer interface {
	Create(userID int64) (token *Token, err error)
	Parse(token string) (*Token, error)
}

type signer struct {
	tokenType  string
	privateKey ed25519.PrivateKey
	expiry     time.Duration
	issuer     string
}

var (
	err   error
	b     []byte
	block *pem.Block
	pub   ed25519.PublicKey
	priv  ed25519.PrivateKey
)

func New() (Signer, error) {

	privateKey, err := jwt.ParseEdPrivateKeyFromPEM([]byte(`-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIO6Jel91mzd++JsBxk/X3sdQxxkqYNnK7BOUPk1eANYF
-----END PRIVATE KEY-----`))
	if err != nil {
		return nil, err
	}
	expiry := time.Duration(86000)

	return &signer{
		privateKey: privateKey.(ed25519.PrivateKey),
		expiry:     expiry,
		issuer:     "MyId",
	}, nil
}

func (t *signer) Create(userID int64) (*Token, error) {
	now := time.Now()
	iat := now.UTC().Unix()
	exp := now.Add(t.expiry * time.Second)
	id := uuid.New().String()
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Id:        id,
			ExpiresAt: exp.Unix(),
			IssuedAt:  iat,
			NotBefore: iat,
			Issuer:    t.issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims).SignedString(t.privateKey)

	return &Token{
		Raw:    token,
		Claims: claims,
	}, err
}

func (t *signer) Parse(token string) (*Token, error) {
	tk, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.privateKey.Public(), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(*Claims); ok && tk.Valid {
		return &Token{
			Claims: claims,
			Raw:    token,
		}, nil
	} else {
		return nil, err
	}
}
