package controller

import (
	"encoding/json"
	"fmt"
	"go-native-crud/src/model"
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

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	service.AddBook(book)
	w.WriteHeader(http.StatusCreated)
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book model.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	err = service.UpdateBook(id, book)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update book: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteBook(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete book: %v", err), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
