package main

import (
	"fmt"
	"go-native-crud/src/router"
	"log"
	"net/http"
)

func main() {
	router.SetupBookRoutes()

	port := ":8080"
	fmt.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
