package http

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/internal/http/jwt"
	"github.com/wormi4ok/menuplanner/internal/http/oauth"
	"golang.org/x/text/language"
)

const authTokenDuration = time.Hour
const refreshTokenDuration = 7 * 24 * time.Hour

type TokenGenerator interface {
	CreateAccessToken(user *internal.User, expiresIn time.Duration) (string, error)
	CreateRefreshToken(user *internal.User, expiresIn time.Duration) (string, error)
}

type userEndpoint struct {
	token   TokenGenerator
	storage internal.UserRepository
	oAuth   *oauth.Google
}

// Routes creates a REST router for the user authentication
func (e userEndpoint) Routes() chi.Router {
	rand.Seed(time.Now().UTC().UnixNano())
	r := chi.NewRouter()

	r.Post("/signup", e.Signup())
	r.Post("/login", e.Login())
	if e.oAuth != nil {
		r.Post("/login/google", e.GoogleAuth(e.oAuth))
	}

	return r
}

type authTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func newAuthTokenResponse(accessToken string, refreshToken string) *authTokenResponse {
	return &authTokenResponse{
		AccessToken:  accessToken,
		TokenType:    "bearer",
		ExpiresIn:    int(authTokenDuration.Seconds()),
		RefreshToken: refreshToken,
	}
}

func (e *userEndpoint) Signup() http.HandlerFunc {
	type request struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if existingUser, err := e.storage.ReadUserByEmail(r.Context(), req.Email); existingUser != nil {
			http.Error(w, "User with this email already exists", http.StatusConflict)
			return
		} else if !internal.ErrorIs(err, internal.ErrorNotFound) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := internal.NewUser(req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if locale := getRequestLocale(r); locale != "" {
			user.Locale = locale
		}

		if err := e.storage.CreateUser(r.Context(), user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		at, rt, err := e.tokenPair(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON(w, newAuthTokenResponse(at, rt))
	}
}

func (e *userEndpoint) Get() http.HandlerFunc {
	type response struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := jwt.UserID(r.Context())
		user, err := e.storage.ReadUser(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := response{
			Name:  user.Name,
			Email: user.Email,
		}

		responseJSON(w, res)
	}
}

func (e *userEndpoint) Login() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := e.storage.ReadUserByEmail(r.Context(), req.Email)
		if err != nil {
			http.Error(w, "Login failed; Invalid email or password", http.StatusUnauthorized)
			return
		}
		if !user.HasPassword(req.Password) {
			http.Error(w, "Login failed; Invalid email or password", http.StatusUnauthorized)
			return
		}

		at, rt, err := e.tokenPair(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON(w, newAuthTokenResponse(at, rt))
	}
}

func (e *userEndpoint) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := jwt.UserID(r.Context())

		user, err := e.storage.ReadUser(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		at, rt, err := e.tokenPair(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseJSON(w, newAuthTokenResponse(at, rt))
	}
}

func (e *userEndpoint) GoogleAuth(googleOAuth *oauth.Google) http.HandlerFunc {
	type request struct {
		AuthCode string `json:"code"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userInfo, err := googleOAuth.UserInfo(r.Context(), req.AuthCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusFailedDependency)
			return
		}

		if existingUser, err := e.storage.ReadUserByEmail(r.Context(), userInfo.Email); existingUser != nil {
			at, rt, err := e.tokenPair(existingUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			responseJSON(w, newAuthTokenResponse(at, rt))
		} else if internal.ErrorIs(err, internal.ErrorNotFound) {
			user := &internal.User{
				Name:    userInfo.Name,
				Email:   userInfo.Email,
				Picture: userInfo.Picture,
				Locale:  userInfo.Locale,
			}

			if err := e.storage.CreateUser(r.Context(), user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			at, rt, err := e.tokenPair(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			responseJSON(w, newAuthTokenResponse(at, rt))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (e *userEndpoint) tokenPair(user *internal.User) (accessToken string, refreshToken string, err error) {
	if accessToken, err = e.token.CreateAccessToken(user, authTokenDuration); err != nil {
		return
	}

	if refreshToken, err = e.token.CreateRefreshToken(user, refreshTokenDuration); err != nil {
		return
	}

	return
}

func getRequestLocale(r *http.Request) string {
	t, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if err == nil && t != nil && len(t) > 0 {
		return t[0].String()
	}

	return ""
}
