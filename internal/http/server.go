package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/internal/http/jwt"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func NewServer(
	host string, port int, jwtSecret string,
	oAuth *OAuth,
	recipes internal.RecipeRepository,
	courses internal.CourseReader,
	weeks internal.WeekRepository,
	users internal.UserRepository,
	docs []byte,
) *Server {
	r := router()

	we := weekEndpoint{storage: weeks, filler: internal.NewGapFiller(recipes, courses)}
	ue := userEndpoint{&jwt.Generator{Secret: jwtSecret}, users, oAuth}

	// Public routes
	r.Mount("/auth", ue.Routes())
	r.Handle("/docs*", docsEndpoint{docs})
	r.Get("/health", healthEndpoint)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwt.AccessTokenVerifier(jwtSecret))
		r.Use(jwt.AccessTokenAuthenticator)

		r.Get("/", we.Get())
		r.Get("/user/me", ue.Get())
		r.Mount("/week", we.Routes())
		r.Mount("/recipe", recipeEndpoint{recipes}.Routes())
		r.Mount("/course", courseEndpoint{courses}.Routes())
	})

	r.Group(func(r chi.Router) {
		r.Use(jwt.RefreshTokenVerifier(jwtSecret))
		r.Use(jwt.RefreshTokenAuthenticator(users))

		r.Post("/token/refresh", ue.Refresh())
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%v\n", err)
		}
	}()

	return &Server{srv: srv}
}

func router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Authorization"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	}))
	return r
}

func healthEndpoint(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, "What's on the menu?")
}

func readJSON(r *http.Request, req interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("missing or malformed payload")
	}
	_ = r.Body.Close()

	return json.Unmarshal(body, &req)
}

func responseJSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		_, _ = io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
