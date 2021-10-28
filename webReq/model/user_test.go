package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	//do whatever you want before testing
	m.Run() // this would run TestUser func
}

func TestUser(t *testing.T) {
	fmt.Println("testing User related func...")
	t.Run("testing AddUser with prepare:", testAddUser)
	t.Run("testing AddUser without prepare:", testAddUser2)
}

func testAddUser(t *testing.T) {
	fmt.Println("testing user add with prepate...")
	user := &User{}
	user.AddUser()
}
func testAddUser2(t *testing.T) {
	fmt.Println("testing user add without prepate...")
	user := &User{}
	user.AddUser2()
}

/*func TestUser_AddUser(t *testing.T) {
	fmt.Println("testing user add with prepate...")
	user := &User{}
	user.AddUser()
}

func TestUser_AddUser2(t *testing.T) {
	fmt.Println("testing user add without prepate...")
	user := &User{}
	user.AddUser2()
}*/
func TestUser_GetUserById(t *testing.T) {
	fmt.Println("select one row...")
	user := &User{ID: 1}
	u, _ := user.GetUserById()
	fmt.Println(u.Email)
}
func TestUser_GetUsers(t *testing.T) {
	fmt.Println("select all row...")
	user := &User{}
	us, _ := user.GetUsers()
	for _, v := range us {
		fmt.Printf("<%v, %v>\n", v.ID, v.Username)
	}
}
