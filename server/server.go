package server

import (
	"github.com/binarymodder/ServerBTI"
	"github.com/binarymodder/ServerBTI/graph"
	"github.com/binarymodder/ServerBTI/src/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{},
		},
	))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("🚀 GraphQL server running at http://localhost:%s/ ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
