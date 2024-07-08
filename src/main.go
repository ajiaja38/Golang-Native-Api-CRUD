package main

import (
	"go-native-crud/src/router"
	"log"
	"net/http"
	"os"
)

func main() {
	router.SetupBookRoutes()

	port := ":8080"

	logger := log.New(os.Stdout, "API", log.LstdFlags)
	logger.Printf("Server started on port %s", port)
	logger.Fatal(http.ListenAndServe(port, nil))
}
