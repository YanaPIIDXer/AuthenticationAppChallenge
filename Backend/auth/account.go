package auth

import (
	"database/sql"
	"fmt"

	"api/register/result_code"
	"db"
)

// 基本認証登録
func RegisterBasicAuth(email string, password string) int {
	if email == "" || password == "" { return result_code.EmptyParam }

	var resultCode = result_code.RegisterSuccess
	err := msqldrv.Access(func(db *sql.DB) {		
		var dummy = 0
		err := db.QueryRow("SELECT id FROM basic_auth where email=?", email).Scan(&dummy)
		if err == nil {
			resultCode = result_code.UsedEmail
			return
		} else if err != sql.ErrNoRows {
			fmt.Println(err.Error())
			resultCode = result_code.Fatal
			return
		}

		tx, _ := db.Begin()
		stmt, err := db.Prepare("INSERT INTO accounts VALUES()")
		result, err := stmt.Exec()
		if err != nil {
			fmt.Println(err.Error())
			resultCode = result_code.Fatal
			tx.Rollback()
			return
		}

		id, _ := result.LastInsertId()
		err = db.QueryRow("INSERT INTO basic_auth VALUES(?, ?, ?)", id, email, password).Scan()
		if err != sql.ErrNoRows {
			fmt.Println(err.Error())
			resultCode = result_code.Fatal
			tx.Rollback()
			return
		}
		tx.Commit()
	})

	if err != nil {
		resultCode = result_code.Fatal
		fmt.Println(err.Error())
	}
	
	return resultCode
}
