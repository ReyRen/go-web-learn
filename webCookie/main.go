package main

import (
	"fmt"
	"net/http"
)

//setCookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "user",
		Value: "admin",
	}
	/*w.Header().Set("Set-Cookie", cookie.String())
	w.Header().Add("Set-Cookie", cookie.String())*/

	http.SetCookie(w, &cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	//	cookies := r.Header["Cookie"]
	cookies, _ := r.Cookie("user")
	fmt.Println(cookies)
}

func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookie", getCookie)
	http.ListenAndServe(":8080", nil)
}
