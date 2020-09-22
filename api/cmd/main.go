package main

import (
	"context"
	"net/http"
	"os"

	"github.com/awesomebusiness/uinvest/config"
	"github.com/awesomebusiness/uinvest/internal/generated"
	"github.com/awesomebusiness/uinvest/internal/resolver"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/repository"
	"github.com/awesomebusiness/uinvest/internal/service/authentication/usecase"
	"github.com/awesomebusiness/uinvest/pkg/twillio"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	ctx := context.Background()

	// load environment variable from .env
	_ = godotenv.Load()

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true

	log.SetFormatter(customFormatter)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configInternal, err := config.NewConfigMap()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer configInternal.Database.Client.Close()

	if err := configInternal.Database.Client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	accountID := os.Getenv("TWILLIO_ACCOUNT_ID")
	authToken := os.Getenv("TWILLIO_AUTH_TOKEN")
	twillioPhoneNumber := os.Getenv("TWILLIO_PHONE_NUMBER")

	twillioClient, err := twillio.NewTwillioClient(accountID, authToken, twillioPhoneNumber)
	if err != nil {
		log.Fatalf("failed create twillio client : %v", err)
	}

	authrepo := repository.NewAuthenticationRepository(configInternal.Database.Client, twillioClient)
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
