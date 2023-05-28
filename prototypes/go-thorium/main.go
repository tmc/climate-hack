package main

import (
	"context"
	"flag"
	"fmt"
	"go-thorium/graph"
	"go-thorium/thoriumfacts"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"github.com/tmc/langchaingo/llms/openai"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func main() {
	flag.Parse()
	// set up logging
	logger, err := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync() //nolint:errcheck
	if err != nil {
		log.Println("can't initialize zap logger:", err)
	}

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	llm, err := openai.New()
	if err != nil {
		return fmt.Errorf("failed to create LLM: %w", err)
	}
	redisClient, err := newRedisClient()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	service := thoriumfacts.NewService(llm, redisClient)
	resolver := graph.NewResolver(service)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	mux := http.NewServeMux()

	mux.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	mux.Handle("/graphql", srv)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	handler := cors.Default().Handler(mux)
	return http.ListenAndServe(":"+port, handler)
}

func newRedisClient() (*redis.Client, error) {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
		log.Println("redis url not set, using default of", redisURL)
	}
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis url: %w", err)
	}
	c := redis.NewClient(opts)
	_, err = c.Ping(context.Background()).Result()
	return c, err
}
