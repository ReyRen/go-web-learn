package dao

import (
	"go-web-learn/bookstore/model"
	"go-web-learn/bookstore/utils"
)

//AddSession
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionById
func GetSessionById(sessID string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessID)
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}
