package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:Hh!067067@tcp(localhost:3306)/bookstore") // root:admin123@/test
	if err != nil {
		panic(err.Error())
	}
}
