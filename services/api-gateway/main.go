package main

import (
	"github.com/imroc/req/v3"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/generated"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/resolver"
)

const defaultPort = "8080"

func main() {

	req.DevMode()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
