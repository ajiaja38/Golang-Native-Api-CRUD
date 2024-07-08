package controller

import (
	"encoding/json"
	"fmt"
	"go-native-crud/src/service"
	"net/http"
	"strconv"
)

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := service.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	book, err := service.GetBookByID(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("Book not found: %v", err), http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(book)
}
