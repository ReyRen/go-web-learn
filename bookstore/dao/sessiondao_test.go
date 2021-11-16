package dao

import (
	"fmt"
	"go-web-learn/bookstore/model"
	"testing"
)

func TestSession(t *testing.T) {
	fmt.Println("test sessiondao func...")
	//t.Run("add session test:", testAddSession)
	//t.Run("delete session test:", testDeleteSession)
	t.Run("get session test:", testGetSessionById)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838385438",
		UserName:  "fucker",
		UserID:    2,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838385438")
}

func testGetSessionById(t *testing.T) {
	DeleteSession("13838385438")
}
