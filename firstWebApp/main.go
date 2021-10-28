package main

import (
	"fmt"
	"net/http"
)

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world", r.URL.Path) // Hello world /
}

func main() {
	/*
		HandleFunc is a adapter, it's can convert normal func to a HTTP handler
	*/
	http.HandleFunc("/", handler)
	//route
	http.ListenAndServe(":8080", nil) // nil would be use DefaultServeMux

	/*
		流程是：请求到达后，会先调用多路复用器（nil则使用默认的），
		多路复用器就会根据http.HandleFunc中已经注册的规则，进行handler的调用
	*/
}
