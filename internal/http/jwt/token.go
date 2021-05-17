package jwt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func AccessTokenVerifier(jwtSecret string) func(http.Handler) http.Handler {
	tokenAuth := jwtauth.New("HS512", []byte(jwtSecret), nil)
	return jwtauth.Verifier(tokenAuth)
}

func AccessTokenAuthenticator(next http.Handler) http.Handler {
	return jwtauth.Authenticator(next)
}

func RefreshTokenVerifier(jwtSecret string) func(http.Handler) http.Handler {
	tokenAuth := jwtauth.New("HS512", []byte(jwtSecret), nil)
	return func(next http.Handler) http.Handler {
		return jwtauth.Verify(tokenAuth, TokenFromBody, jwtauth.TokenFromHeader)(next)
	}
}

func RefreshTokenAuthenticator(reader internal.UserReader) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			id, err := strconv.Atoi(token.Subject())
			if err != nil {
				http.Error(w, "malformed token", http.StatusNotAcceptable)
				return
			}

			u, err := reader.ReadUser(r.Context(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusPreconditionFailed)
				return
			}

			if token == nil || jwt.Validate(token, jwt.WithClaimValue("key", u.Key)) != nil {
				http.Error(w, http.StatusText(401), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func TokenFromBody(r *http.Request) string {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	b := bytes.NewBuffer(make([]byte, 0))
	reader := io.TeeReader(r.Body, b)

	if err := json.NewDecoder(reader).Decode(&req); err != nil {
		return ""
	}

	r.Body = io.NopCloser(b)
	return req.RefreshToken
}
