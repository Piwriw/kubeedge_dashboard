package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"new-ec-dashboard/models"
)
const secret = "Pirwiw.com"


// encryptPassword : 加密密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	// password is T
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
