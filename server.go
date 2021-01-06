package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davidyap2002/user-go/api"
	"github.com/davidyap2002/user-go/directives"
	"github.com/davidyap2002/user-go/graph"
	"github.com/davidyap2002/user-go/graph/generated"
	"github.com/davidyap2002/user-go/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {

	_, err := api.CheckEmail(context.Background(), "davidyasagasgasgp11les@gmail.com")

	fmt.Println(err)
	// migration.MigrateTable()	
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middleware.Middleware())

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.IsLogin = directives.IsLogin

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
