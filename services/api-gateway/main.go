package main

import (
	"github.com/imroc/req/v3"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/auth"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/generated"
	"gitlab.informatik.hs-augsburg.de/flomon/waaf/services/api-gateway/graph/resolver"

	"github.com/go-chi/chi"
)

const defaultPort = "10003"

func main() {

	req.DevMode()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
