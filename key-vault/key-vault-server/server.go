package main

import (
	"key-vault-server/graphql"
	"key-vault-server/graphql/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3030"

func main() {
	port := os.Getenv("PORT")
	url := os.Getenv("ROOT_URL")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to %s:%s/ for GraphQL playground", url, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}