package main

import (
	log "log"
	http "net/http"
	os "os"

	"github.com/gorilla/mux"

	handler "github.com/99designs/gqlgen/handler"
	gqlgen_dataloader "github.com/aneri/gqlgen-dataloader"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := mux.NewRouter()
	router.Use(gqlgen_dataloader.Middleware)
	router.Use(gqlgen_dataloader.ApplicationLoaderMiddleware)
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", handler.GraphQL(gqlgen_dataloader.NewExecutableSchema(gqlgen_dataloader.Config{Resolvers: &gqlgen_dataloader.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
