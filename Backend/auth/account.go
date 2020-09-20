package auth

import (
	"database/sql"

	"api/register/result_code"
	"db"
)

// 基本認証登録
func RegisterBasicAuth(email string, password string) int {
	if email == "" || password == "" { return result_code.EmptyParam }

	var resultCode = result_code.RegisterSuccess
	msqldrv.Access(func(db *sql.DB) {		
		err := db.QueryRow("SELECT id FROM basic_auth where email=?", 1).Scan(&email)
		if err == nil {
			resultCode = result_code.UsedEmail
		} else if err != sql.ErrNoRows {
			resultCode = result_code.Fatal
		}
	})
	
	return resultCode
}
