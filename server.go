package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/99designs/gqlgen/handler"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"grapgQL/database"
	"grapgQL/dbhelper"
	"grapgQL/graph"
	_ "grapgQL/graph"
	"grapgQL/middleware"
	"log"
	_ "log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {

	db, err := database.DbConnection()
	if err != nil {
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create the GraphQL handler with the resolver
	gqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			DAO: dbhelper.DAO{
				DB: db,
			},
		}}))

	// Create a new router
	router := mux.NewRouter()

	// Create a separate subRouter for API routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	// Apply the API session key middleware to the API routes
	apiRouter.Use(middleware.AuthMiddleware)
	apiRouter.Use(middleware.CorsMiddleware)

	// Register the GraphQL API route under the /api endpoint
	apiRouter.Handle("/graphql", gqlHandler)
	router.Handle("/", playground.Handler("GraphQL playground", "/api/graphql"))

	// Define the login route outside the API subRouter
	router.HandleFunc("/login", graph.LoginHandler)

	// Create a CORS handler to allow cross-origin requests
	corsHandler := cors.Default().Handler(router)

	// Create a server with the CORS handler
	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler,
	}

	// Start the server
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(server.ListenAndServe())
}
