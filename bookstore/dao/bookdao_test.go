package dao

import (
	"fmt"
	"go-web-learn/bookstore/model"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("test bookdao func...")
	//t.Run("Get all books:", testGetBooks)
	//t.Run("Add book:", testAddBooks)
	//t.Run("Delete book:", testDeleteBook)
	//t.Run("Get book:", testGetBookById)
	//t.Run("Update book:", testUpdateBook)
	//t.Run("Get page books:", testGetPageBooks)
	t.Run("Get page books by price:", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("No.%v book's info:%v\n", k+1, v)
	}
}

func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:     "testTitle",
		Author:    "testAuthor",
		Price:     100,
		Sales:     100,
		Stock:     100,
		ImagePath: "/static/img/default.jpg",
	}
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	DeleteBook("34")
}

func testGetBookById(t *testing.T) {
	book, _ := GetBookById("2")
	fmt.Println(book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:     1,
		Title:  "解忧杂货店2",
		Author: "东野圭吾2",
		Price:  102,
		Sales:  102,
		Stock:  102,
	}
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("2")
	fmt.Println(page.Books[0].Title)
}
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("1", "10", "30")
	fmt.Println(page.Books[0].Title)
}
