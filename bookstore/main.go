package main

import (
	"go-web-learn/bookstore/controller"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore/views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore/views/pages/"))))
	http.HandleFunc("/entry", controller.IndexHandler)
	http.HandleFunc("/login", controller.Login) //handler func(ResponseWriter, *Request)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName) // Ajax
	//http.HandleFunc("/getBooks", controller.GetBooks)
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	//http.HandleFunc("/addBook", controller.AddBooks)
	http.HandleFunc("/deleteBook", controller.DeleteBooks)
	http.HandleFunc("/toUpdateBookPage", controller.UpdateBookPage)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)

	http.ListenAndServe(":8080", nil)
}
