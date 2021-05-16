package jwt

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/wormi4ok/menuplanner/internal"
)

type Generator struct {
	Secret string
}

func (g *Generator) CreateAccessToken(user *internal.User, expiresIn time.Duration) (string, error) {
	t := jwt.New()
	_ = t.Set(jwt.NotBeforeKey, time.Now().Add(-time.Second))
	_ = t.Set(jwt.ExpirationKey, time.Now().Add(expiresIn))
	_ = t.Set(jwt.SubjectKey, fmt.Sprintf("%d", user.ID))

	signedToken, err := jwt.Sign(t, jwa.HS512, []byte(g.Secret))

	return string(signedToken), err
}

func (g *Generator) CreateRefreshToken(user *internal.User, expiresIn time.Duration) (string, error) {
	t := jwt.New()
	_ = t.Set(jwt.NotBeforeKey, time.Now().Add(-time.Second))
	_ = t.Set(jwt.ExpirationKey, time.Now().Add(expiresIn))
	_ = t.Set(jwt.SubjectKey, fmt.Sprintf("%d", user.ID))
	_ = t.Set("key", user.Key)

	signedToken, err := jwt.Sign(t, jwa.HS512, []byte(g.Secret))

	return string(signedToken), err
}

func UserID(ctx context.Context) int {
	token, _ := ctx.Value(jwtauth.TokenCtxKey).(jwt.Token)
	id, _ := strconv.Atoi(token.Subject())

	return id
}

func Verifier(jwtSecret string) func(http.Handler) http.Handler {
	tokenAuth := jwtauth.New("HS512", []byte(jwtSecret), nil)
	return jwtauth.Verifier(tokenAuth)
}

func Authenticator(next http.Handler) http.Handler {
	return jwtauth.Authenticator(next)
}
