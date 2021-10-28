package main

import (
	"go-web-learn/bookstore/controller"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("bookstore/views/index.html"))
	t.Execute(w, "")
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("bookstore/views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("bookstore/views/pages/"))))
	http.HandleFunc("/entry", IndexHandler)
	http.HandleFunc("/login", controller.Login) //handler func(ResponseWriter, *Request)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName) // Ajax

	http.ListenAndServe(":8080", nil)
}
