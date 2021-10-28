package main

import (
	"encoding/json"
	"fmt"
	"go-web-learn/webReq/model"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "send request url is:", r.URL.Path)
	fmt.Fprintln(w, "request strings after url:", r.URL.RawQuery)
	fmt.Fprintln(w, "request Header all info:", r.Header)
	fmt.Fprintln(w, "request Header Accept-Encoding info:", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "request Header Accept-Encoding paramater:", r.Header.Get("Accept-Encoding"))
	// get body length
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body) // already read all body info to buf, nothing left in r

	fmt.Fprintln(w, "request body info:", string(body))

	/*//1. Parse form
	r.ParseForm()
	//2. Get request paramater
	fmt.Fprintln(w, "Request action url Paramater:", r.Form) // // get paramaters from actoin URL
	fmt.Fprintln(w, "Request post Paramater:", r.PostForm) // // get paramaters from post*/

	//2. FormValue/PostFormValu
	fmt.Fprintln(w, "Form table username:", r.PostFormValue("username"))
	fmt.Fprintln(w, "URL User:", r.FormValue("user"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// create User
	user := model.User{
		ID:       1,
		Username: "zhangsan",
		Password: "123456",
		Email:    "zhangsan@ia.ac.cn",
	}
	js, _ := json.Marshal(user)
	w.Write(js)
}

func testRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedirect)
	http.ListenAndServe(":8080", nil)
}
