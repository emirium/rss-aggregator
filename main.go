package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// routing, middleware
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: false,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthcheck", handlerHealthcheck)
	apiRouter.Get("/error", handlerError)
	router.Mount("/api", apiRouter)

	// server setup
	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	log.Printf("Server started on port: %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}