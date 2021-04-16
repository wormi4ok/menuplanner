package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/wormi4ok/menuplanner/internal"
)

type Server struct {
	srv *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func NewServer(
	host string, port int,
	recipes internal.RecipeRepository,
	weeks internal.WeekRepository,
	docs []byte,
) *Server {
	r := router()

	we := weekEndpoint{storage: weeks, filler: internal.NewGapFiller(recipes)}
	r.Get("/", we.Get())
	r.Mount("/week", we.Routes())
	r.Mount("/recipe", recipeEndpoint{recipes}.Routes())
	r.Handle("/docs*", docsEndpoint{docs})

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
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return r
}
