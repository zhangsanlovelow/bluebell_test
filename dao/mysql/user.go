package mysql

import (
	"bullbell_test/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
)

// CheckUserExist check if user exist in database
// return nil if user exist, otherwise return ErrorUserExist
func CheckUserExist(username string) error {
	sqlStr := `select count(username) from user where username =?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// InsertUser insert user into database
// return nil if insert success, otherwise return error
func InsertUser(user *models.User) error {
	sqlStr := `insert into user(user_id, username, password, email) values(?,?,?,?)`
	user.Password = encryptPassword(user.Password)
	if _, err := db.Exec(sqlStr, user.UserID, user.UserName, user.Password, user.Email); err != nil {
		return err
	}
	return nil
}

// UpdateUser enhance userpassword
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte("secrect" + oPassword))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))

}

// Login login user
// return nil if login success, otherwise return error
func Login(user *models.User) error {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username =?`
	if err := db.Get(user, sqlStr, user.UserName); err != nil {
		if err == sql.ErrNoRows {

			return ErrorUserNotExist
		}
		fmt.Println(err)
		return err
	}
	if encryptPassword(oPassword) != user.Password {

		return ErrorInvalidPassword
	}

	return nil

}
