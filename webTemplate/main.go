package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 1. parse template
	t, err := template.ParseFiles("/home/yuanren/go/src/go-web-learn/webTemplate/index.html")
	if err != nil {
		fmt.Println("template.ParseFiles err=", err)
	}
	t.Execute(w, "Hello template")
}

func main() {

	http.HandleFunc("/testTemplate", testTemplate)
	http.ListenAndServe(":8080", nil)
}
