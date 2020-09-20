package auth

import (
	"database/sql"
	"fmt"
    "golang.org/x/crypto/bcrypt"

	"api/login/login_result_code"
	"api/register/register_result_code"
	"db"
)

// 基本認証でログイン
func LoginWithBasicAuth(email string, password string) int {
	var resultCode = login_result_code.LoginSuccess
	err := msqldrv.Access(func(db *sql.DB) {
		var hash = ""
		err := db.QueryRow("SELECT password FROM basic_auth where email=?", email).Scan(&hash)
		if err != nil {
			if err == sql.ErrNoRows {
				resultCode = login_result_code.NotRegister
			} else {
				fmt.Println(err.Error())
				resultCode = login_result_code.Fatal
			}
		}
		if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
			resultCode = login_result_code.NotRegister
		}
	})
	if err != nil { return login_result_code.Fatal }
	
	return resultCode
}

// 基本認証登録
func RegisterBasicAuth(email string, password string) int {
	if email == "" || password == "" { return register_result_code.EmptyParam }

	password = getPasswordHash(password)
	var resultCode = register_result_code.RegisterSuccess
	err := msqldrv.Access(func(db *sql.DB) {
		var dummy = 0
		err := db.QueryRow("SELECT id FROM basic_auth where email=?", email).Scan(&dummy)
		if err == nil {
			resultCode = register_result_code.UsedEmail
			return
		} else if err != sql.ErrNoRows {
			fmt.Println(err.Error())
			resultCode = register_result_code.Fatal
			return
		}

		tx, _ := db.Begin()
		stmt, err := db.Prepare("INSERT INTO accounts VALUES()")
		defer stmt.Close()
		result, err := stmt.Exec()
		if err != nil {
			fmt.Println(err.Error())
			resultCode = register_result_code.Fatal
			tx.Rollback()
			return
		}

		id, _ := result.LastInsertId()
		err = db.QueryRow("INSERT INTO basic_auth VALUES(?, ?, ?)", id, email, password).Scan()
		if err != sql.ErrNoRows {
			fmt.Println(err.Error())
			resultCode = register_result_code.Fatal
			tx.Rollback()
			return
		}
		tx.Commit()
	})

	if err != nil {
		resultCode = register_result_code.Fatal
		fmt.Println(err.Error())
	}
	
	return resultCode
}

// パスワードのハッシュ値を取得
func getPasswordHash(plain string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(hash)
}
