package http

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/internal/http/jwt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const authTokenDuration = time.Hour
const refreshTokenDuration = 7 * 24 * time.Hour
const userInfoURL = "https://openidconnect.googleapis.com/v1/userinfo"

type TokenGenerator interface {
	CreateAccessToken(user *internal.User, expiresIn time.Duration) (string, error)
	CreateRefreshToken(user *internal.User, expiresIn time.Duration) (string, error)
}

type OAuth struct {
	ClientID     string
	ClientSecret string
}

type userEndpoint struct {
	token   TokenGenerator
	storage internal.UserRepository
	oAuth   *OAuth
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

		if err := e.storage.CreateUser(r.Context(), user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		at, rt, err := e.tokenPair(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := authTokenResponse{
			AccessToken:  at,
			TokenType:    "bearer",
			ExpiresIn:    int(authTokenDuration.Seconds()),
			RefreshToken: rt,
		}

		responseJSON(w, res)
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

		res := authTokenResponse{
			AccessToken:  at,
			TokenType:    "bearer",
			ExpiresIn:    int(authTokenDuration.Seconds()),
			RefreshToken: rt,
		}

		responseJSON(w, res)
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

		res := authTokenResponse{
			AccessToken:  at,
			TokenType:    "bearer",
			ExpiresIn:    int(authTokenDuration.Seconds()),
			RefreshToken: rt,
		}

		responseJSON(w, res)
	}
}

func (e *userEndpoint) GoogleAuth(config *OAuth) http.HandlerFunc {
	type request struct {
		AuthCode string `json:"code"`
	}

	type googleResponse struct {
		Name          string `json:"name"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Locale        string `json:"locale"`
	}

	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  "postmessage",
		Endpoint:     google.Endpoint,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req      request
			userInfo googleResponse
		)

		if err := readJSON(r, &req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := conf.Exchange(r.Context(), req.AuthCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusFailedDependency)
			return
		}

		log.Printf("%+v", token.AccessToken)

		authClient := conf.Client(r.Context(), token)

		res, err := authClient.Get(userInfoURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusFailedDependency)
			return
		}

		if body, err := io.ReadAll(res.Body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			_ = r.Body.Close()
			if err = json.Unmarshal(body, &userInfo); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if existingUser, err := e.storage.ReadUserByEmail(r.Context(), userInfo.Email); existingUser != nil {
			at, rt, err := e.tokenPair(existingUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res := authTokenResponse{
				AccessToken:  at,
				TokenType:    "bearer",
				ExpiresIn:    int(authTokenDuration.Seconds()),
				RefreshToken: rt,
			}

			responseJSON(w, res)
		} else if internal.ErrorIs(err, internal.ErrorNotFound) {
			user := &internal.User{
				Name:  userInfo.Name,
				Email: userInfo.Email,
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

			res := authTokenResponse{
				AccessToken:  at,
				TokenType:    "bearer",
				ExpiresIn:    int(authTokenDuration.Seconds()),
				RefreshToken: rt,
			}

			responseJSON(w, res)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (e *userEndpoint) tokenPair(user *internal.User) (string, string, error) {
	at, err := e.token.CreateAccessToken(user, authTokenDuration)
	if err != nil {
		return "", "", err
	}

	rt, err := e.token.CreateRefreshToken(user, refreshTokenDuration)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}
