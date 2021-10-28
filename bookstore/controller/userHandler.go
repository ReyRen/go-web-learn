package controller

import (
	"go-web-learn/bookstore/dao"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//get username and password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {
		// correct username and password
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/login.html"))
		t.Execute(w, "username/password incorrect!")
	}
}

func Regist(w http.ResponseWriter, r *http.Request) {
	//get username and password
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		// correct username and password
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/regist.html"))
		t.Execute(w, "username already exist!")
	} else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("bookstore/views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//get username and password
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("username already exist"))
	} else {
		w.Write([]byte("<font style='color:green'>username pass</font>"))
	}
}
