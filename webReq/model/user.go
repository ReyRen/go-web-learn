package model

import (
	"fmt"
	"go-web-learn/webDB/utils"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) AddUser() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("utils.Db.Prepare err=", err)
		return err
	}
	_, err = inStmt.Exec("admin", "123456", "admin@guigu.com")
	if err != nil {
		fmt.Println("inStmt.Exec err=", err)
		return err
	}
	return nil
}

// not include prepare
func (user *User) AddUser2() error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, "admin2", "123456", "admin2@guigu.com")
	if err != nil {
		fmt.Println("utils.Db.Exec err=", err)
		return err
	}

	return nil
}

func (user *User) GetUserById() (*User, error) {
	sqlStr := "select id,username,password,email from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, user.ID)
	var id int
	var username string
	var password string
	var email string

	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		fmt.Println("row.Scan err=", err)
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

func (user *User) GetUsers() ([]*User, error) {
	sqlStr := "select id,username,password,email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("utils.Db.Query err=", err)
		return nil, err
	}

	// create slice
	var users []*User

	for rows.Next() {
		var id int
		var username string
		var password string
		var email string

		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			fmt.Println("row.Scan err=", err)
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}

	return users, nil
}
