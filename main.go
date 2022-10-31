// main package
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"38_gitactions/handler"
)

func main() {
	router := httprouter.New()

	router.GET("/health", handler.HealthCheck)
	router.POST("/add", handler.Add)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter)) //nolint:gosec
}
