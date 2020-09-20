package msqldrv

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DBアクセス
func Access(callback func(db *sql.DB)) (bool, error) {
	db, err := sql.Open("mysql", "root:root/auth_app")
	if err != nil { return false, err }
	defer db.Close()

	callback(db)

	return true, nil
}
