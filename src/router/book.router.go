package router

import (
	"go-native-crud/src/controller"
	"net/http"
)

func SetupBookRoutes() {
	http.HandleFunc("/books", controller.GetAllBooksHandler)
	http.HandleFunc("/book", controller.GetBookByIDHandler)
}
