package service

import (
	"fmt"
	"go-native-crud/src/model"
	"sync"
)

var (
	books  []model.Book
	lastID int
	mutex  = &sync.Mutex{}
)

func GetAllBooks() []model.Book {
	return books
}

func GetBookByID(id int) (*model.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}

	return nil, fmt.Errorf("book not found")
}

func AddBook(book model.Book) {
	mutex.Lock()
	defer mutex.Unlock()

	lastID++
	book.ID = lastID
	books = append(books, book)
}

func UpdateBook(id int, book model.Book) error {
	for i := range books {
		if books[i].ID == id {
			mutex.Lock()
			defer mutex.Unlock()

			books[i] = book
			return nil
		}
	}

	return fmt.Errorf("book not found")
}

func DeleteBook(id int) error {
	for i, b := range books {
		if b.ID == id {
			mutex.Lock()
			defer mutex.Unlock()

			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book not found")
}
