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

	"github.com/sethvargo/go-envconfig"
	"github.com/wormi4ok/menuplanner/internal"
	"github.com/wormi4ok/menuplanner/internal/http"
	"github.com/wormi4ok/menuplanner/internal/http/oauth"
	"github.com/wormi4ok/menuplanner/internal/storage"
	"github.com/wormi4ok/menuplanner/internal/storage/mock"
)

type Config struct {
	Host        string `env:"HOST,default=localhost"`
	Port        int    `env:"PORT,default=8081"`
	RecipesJSON string `env:"RECIPES_JSON"`
	WeekJSON    string `env:"WEEK_JSON"`

	MysqlDSN string `env:"MYSQL_DSN,required"`

	JWTSecret string `env:"JWT_SECRET,required"`

	// Oauth2 credentials
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
}

//go:embed docs/index.html
var docs []byte

func main() {
	var c Config
	l := envconfig.PrefixLookuper("MP_", envconfig.OsLookuper())
	if err := envconfig.ProcessWith(context.Background(), &c, l); err != nil {
		panic(err)
	}

	var (
		weekStorage   internal.WeekRepository
		recipeStorage internal.RecipeRepository
		courseStorage internal.CourseReader
		userStorage   internal.UserRepository
	)

	oAuth := loadOAuth(c)

	if c.MysqlDSN != "" {
		db := loadDB(&c)
		weekStorage, recipeStorage, courseStorage, userStorage = db, db, db, db
	} else {
		recipeStorage, weekStorage = loadMocks(&c)
	}

	srv := http.NewServer(c.Host, c.Port, c.JWTSecret, oAuth, recipeStorage, courseStorage, weekStorage, userStorage, docs)
	log.Printf("Starting service on %s port %d...\n", c.Host, c.Port)

	handleServerShutdown(srv)
}

func loadOAuth(c Config) *oauth.Google {
	if c.ClientID == "" || c.ClientSecret == "" {
		return nil
	}
	log.Println("OAuth authentication configured...")
	return &oauth.Google{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
	}
}

func loadDB(c *Config) *storage.DB {
	log.Printf("Connecting to the database...")
	db, err := storage.InitDB(c.MysqlDSN)
	if err != nil {
		log.Printf("Failed to connect: %s", err)
		os.Exit(1)
	}
	return db
}

func loadMocks(c *Config) (*mock.Recipes, *mock.Weeks) {
	log.Println("Using mock storage...")
	mr := &mock.Recipes{}
	if c.RecipesJSON != "" {
		if err := mr.LoadFromFile(c.RecipesJSON); err != nil {
			log.Println("Failed to load recipes from file")
			os.Exit(1)
		}
	}
	wr := &mock.Weeks{Recipes: mr}
	if c.WeekJSON != "" {
		if err := wr.LoadFromFile(c.WeekJSON); err != nil {
			log.Println("Failed to load recipes from file")
			os.Exit(1)
		}
	}
	return mr, wr
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
