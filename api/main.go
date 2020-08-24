package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/awesomebusiness/uinvest/config"
	"github.com/awesomebusiness/uinvest/internal/generated"
	"github.com/awesomebusiness/uinvest/internal/resolver"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/repository"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/usecase"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()

	// load environment variable from .env
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer db.Client.Close()

	if err := db.Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	authrepo := repository.NewAuthenticationRepository(db.Client)
	authuc := usecase.NewAuthenticationUsecase(authrepo)

	config := generated.Config{
		Resolvers: &resolver.Resolver{
			AuthenticationUC: authuc,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
