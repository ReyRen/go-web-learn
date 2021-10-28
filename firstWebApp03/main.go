package main

import (
	"fmt"
	"net/http"
)

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world3", r.URL.Path)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	//route
	http.ListenAndServe(":8080", mux)

}
