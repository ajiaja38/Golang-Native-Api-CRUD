package router

import (
	"go-native-crud/src/controller"
	"net/http"
)

func SetupBookRoutes() {
	http.HandleFunc("/books", controller.GetAllBooksHandler)
	http.HandleFunc("/book", controller.GetBookByIDHandler)
	http.HandleFunc("/add-book", controller.AddBookHandler)
	http.HandleFunc("/update-book", controller.UpdateBookHandler)
	http.HandleFunc("/delete-book", controller.DeleteBookHandler)
}
