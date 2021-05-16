package jwt

import (
	"fmt"
	"time"

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
