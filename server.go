// server.go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pratyakshroy47/gql-go/graph"
	"github.com/pratyakshroy47/gql-go/logger"
	"github.com/pratyakshroy47/gql-go/mongo"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	logger.Init(config.LogLevel)

	db, err := mongo.NewClient(config.MongoURI)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", err)
	}
	defer db.Disconnect()

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle(config.PlaygroundPath, playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	logger.Infof("Server listening on port %s", config.Port)
	logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil))
}
