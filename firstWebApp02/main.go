package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world2", r.URL.Path)
}

func main() {

	myHandler := MyHandler{}

	/*http.Handle("/", &myHandler)

	http.ListenAndServe(":8080", nil)*/

	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		TLSConfig:   nil,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
