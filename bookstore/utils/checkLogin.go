package utils

import (
	"go-web-learn/bookstore/dao"
	"net/http"
)

func IsLogin(r *http.Request) (bool, string) {
	//get cookie by name
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		session, _ := dao.GetSessionById(cookieValue)
		if session.UserID > 0 {
			return true, session.UserName
		}
	}
	return false, ""
}
