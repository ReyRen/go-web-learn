package controller

import (
	"go-web-learn/bookstore/dao"
	"go-web-learn/bookstore/model"
	"go-web-learn/bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
)

/*func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}*/

func AddBooks(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)

	book := &model.Book{
		Title:     r.PostFormValue("title"),
		Author:    r.PostFormValue("author"),
		Price:     price,
		Sales:     int(sales),
		Stock:     int(stock),
		ImagePath: "/static/img/default.jpg",
	}
	dao.AddBook(book)

	GetPageBooks(w, r)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	booID := r.FormValue("bookId")
	dao.DeleteBook(booID)
	GetPageBooks(w, r)
}

func UpdateBookPage(w http.ResponseWriter, r *http.Request) {
	booID := r.FormValue("bookId")
	book, _ := dao.GetBookById(booID)
	if book.ID > 0 {
		// update book
		t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_edit.html"))
		t.Execute(w, book)
	} else {
		// add book
		t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_edit.html"))
		t.Execute(w, "")
	}
}

func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	sales, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	stock, _ := strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	bookId, _ := strconv.ParseInt(r.PostFormValue("bookId"), 10, 0)

	book := &model.Book{
		ID:     int(bookId),
		Title:  r.PostFormValue("title"),
		Author: r.PostFormValue("author"),
		Price:  price,
		Sales:  int(sales),
		Stock:  int(stock),
	}
	if int(bookId) > 0 {
		dao.UpdateBook(book)
	} else {
		dao.AddBook(book)
	}

	GetPageBooks(w, r)
}

func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}

	page, _ := dao.GetPageBooks(pageNo)

	t := template.Must(template.ParseFiles("bookstore/views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}

	page, _ := dao.GetPageBooks(pageNo)

	t := template.Must(template.ParseFiles("bookstore/views/index.html"))
	t.Execute(w, page)
}

func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	var page *model.Page

	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")

	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {

		if pageNo == "" {
			pageNo = "1"
		}

		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	//get cookie
	/*cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := dao.GetSessionById(cookieValue)
		if session.UserID > 0 {
			// already login
			page.IsLogin = true
			page.Username = session.UserName
		}
	}*/
	flag, userName := utils.IsLogin(r) // true is already login, false is not logged
	if flag {
		// already login
		page.IsLogin = true
		page.Username = userName
	}

	t := template.Must(template.ParseFiles("bookstore/views/index.html"))
	t.Execute(w, page)
}
