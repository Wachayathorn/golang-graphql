package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/99designs/gqlgen"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wachayathorn/golang-graphql/config"
	"github.com/wachayathorn/golang-graphql/graph"
	"github.com/wachayathorn/golang-graphql/graph/generated"
)

const defaultPort = "8080"

func main() {
	// Configure Server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to database
	config.ConnectDatabaseBySQLX()
	config.InitDatabase()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// Route
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
