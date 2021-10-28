package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("test userdao func...")
	t.Run("Validate username and password:", testCheckUserNameAndPassword)
	t.Run("Validate username:", testCheckUserName)
	t.Run("Validate save user:", testSaveUser)
}

func testCheckUserNameAndPassword(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("Get User info:", user)
}
func testCheckUserName(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("Get User info:", user)
}
func testSaveUser(t *testing.T) {
	err := SaveUser("admin2", "123456", "admin2@163.com")
	if err != nil {
		return
	}
}
