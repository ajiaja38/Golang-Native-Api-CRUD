package model

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"autor"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}
