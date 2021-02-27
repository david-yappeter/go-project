package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidyap2002/user-go/dataloader"
	"github.com/davidyap2002/user-go/directives"
	"github.com/davidyap2002/user-go/graph"
	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/service"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// migration.MigrateTable()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(dataloader.Dataloader)
	router.Use(service.Middleware())

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.IsLogin = directives.IsLogin

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/auth/google", service.GoogleOauth2RedirectLink)
	router.HandleFunc("/auth/google/callback", service.GoogleOauth2ParseCallback)
	router.HandleFunc("/verify/email", service.EmailVerification)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
