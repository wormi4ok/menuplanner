package main

import (
	"context"
	"encoding/json"
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
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host string `default:"localhost"`
	Port int    `default:"8081"`
}

func main() {
	c := new(Config)
	envconfig.MustProcess("MP", c)

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

	r.Get("/", weekHandler())

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

func weekHandler() http.HandlerFunc {
	r1 := Recipe{
		ID:       1,
		Name:     "Moroccan Carrot Soup",
		Calories: 300,
		Protein:  60,
		Fat:      50,
		Carbs:    50,
	}

	r2 := Recipe{
		ID:       2,
		Name:     "Ovsyanoblin",
		Calories: 350,
		Protein:  50,
		Fat:      10,
		Carbs:    30,
	}

	r3 := Recipe{
		ID:          3,
		Name:        "Pasta Carbonara",
		Description: "",
		ImageURL:    "",
		Calories:    500,
		Protein:     120,
		Fat:         80,
		Carbs:       260,
	}

	d1 := map[int]Recipe{0: r1, 1: r2}
	d2 := map[int]Recipe{2: r3}
	m := Menu{0: DailyMenu{Recipes: d1}, 1: DailyMenu{d2}}
	week := Week{Menu: m}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(week)
		if err != nil {
			log.Printf("Handler error: %v", err)
			w.WriteHeader(500)
		}
	}
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
