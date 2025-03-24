package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const secret = "qg"

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		err = ErrorUserExist
		return
	}
	return
}

func InsertUser(user *models.User) (err error) {
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	user.Password = encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user *models.User) (err error) {
	sqlStr := "select user_id,username,password from user where username = ?"
	oPassword := user.Password
	err = db.Get(user, sqlStr, user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if encryptPassword(oPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return
}
