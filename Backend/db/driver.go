package msqldrv

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DBアクセス
func Access(callback func(db *sql.DB)) error {
	db, err := sql.Open("mysql", "root:root@tcp(db)/auth_app?parseTime=true&loc=Local&time_zone=%27Asia%2FTokyo%27")
	if err != nil { return err }
	defer db.Close()

	callback(db)

	return nil
}
