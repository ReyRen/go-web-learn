package dao

import (
	"go-web-learn/bookstore/model"
	"go-web-learn/bookstore/utils"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
