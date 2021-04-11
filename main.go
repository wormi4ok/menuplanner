package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/wormi4ok/menuplanner/internal/http"
	"github.com/wormi4ok/menuplanner/internal/storage/mock"
)

type Config struct {
	Host string `default:"localhost"`
	Port int    `default:"8081"`
}

//go:embed docs/index.html
var docs []byte

func main() {
	c := new(Config)
	envconfig.MustProcess("MP", c)

	mr := &mock.Recipes{}
	wr := &mock.Weeks{Recipes: mr}

	srv := http.NewServer(c.Host, c.Port, mr, wr, docs)
	log.Printf("Starting service on %s port %d...\n", c.Host, c.Port)

	handleServerShutdown(srv)
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