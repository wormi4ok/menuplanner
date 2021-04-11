package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/storage/mock"
)

type Config struct {
	Host string `default:"localhost"`
	Port int    `default:"8081"`
}

func main() {
	c := new(Config)
	envconfig.MustProcess("MP", c)

	v := validator.New()
	r := router()
	mr := &mock.Recipes{}
	w := weekEndpoint{
		storage: &mock.Weeks{Recipes: mr},
		filler:  internal.NewGapFiller(mr),
	}

	r.Get("/", w.Get())
	r.Mount("/week", w.Routes())
	r.Mount("/recipe", recipeEndpoint{mr, v}.Routes())
	r.Handle("/docs", http.StripPrefix("/docs", http.FileServer(http.Dir("./docs"))))

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", c.Host, c.Port),
		Handler:      r,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Starting service on %s port %d...\n", c.Host, c.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("%v\n", err)
		}
	}()

	handleServerShutdown(srv)
}

func router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return r
}

func handleServerShutdown(srv *http.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	s := <-signals

	log.Printf("Got %s signal, shutting down server...\n", strings.ToUpper(s.String()))
	// Wait for 5 seconds before shutting down
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Print("Failed to shutdown server gracefully")
		os.Exit(1)
	}
	os.Exit(0)
}
