package auth

import (
	"errors"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"

	"authapp/db"
)

var (
	TokenNotMade = errors.New("Token is not made.")
	TokenIsPeriod = errors.New("Token is period.")
)

// トークンを取得して期限を更新
func GetAndUpdateToken(id int) (string, error) {
	var token = ""
	var accessError error = nil
	err := msqldrv.Access(func(db *sql.DB) {
		var period time.Time
		bFoundRecord, err := getRecord(db, id, &token, &period)
		if err != nil {
			accessError = err
			return
		}

		var current = time.Now()
		if !bFoundRecord {
			// 生成されていない
			accessError = TokenNotMade
			return
		}

		if current.Before(period) {
			// まだ生きているので生存期限のみ更新
			var next = current.Add(12 * time.Hour)
			accessError = db.QueryRow("UPDATE token SET period=? where id=?", next, id).Scan()
			if accessError == sql.ErrNoRows { accessError = nil }
			return
		}
		
		// 寿命
		accessError = TokenIsPeriod
	})
	if err != nil { return "", err }

	return token, accessError
}

// トークン生成
func MakeToken(id int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(string(rand.Int())), bcrypt.DefaultCost)
	if err != nil { return "", err }

	var token = string(hash)
	err = msqldrv.Access(func(db *sql.DB) {
		var dummy = ""
		var dummy2 time.Time
		bFoundRecord, err := getRecord(db, id, &dummy, &dummy2)
		if err != nil { return }

		var period = time.Now().Add(12 * time.Hour)
		if bFoundRecord {
			// レコードは存在するのでUPDATE
			err = db.QueryRow("UPDATE token SET token=?, period=? where id=?", token, period, id).Scan()
			return
		}

		// レコードそのものがないのでINSERT
		err = db.QueryRow("INSERT INTO token VALUES(?, ?, ?)", id, token, period).Scan()
		if err == sql.ErrNoRows { err = nil }		// こいつはRowを返さない。
	})
	
	return token, err
}

// レコード取得
func getRecord(db *sql.DB, id int, token *string, period *time.Time) (bool, error) {
	err := db.QueryRow("SELECT token, period FROM token WHERE id=?", id).Scan(token, period)
	if err != nil {
		if err != sql.ErrNoRows { return false, err }
	}
	var bFound = (err != sql.ErrNoRows)
	err = nil
	return bFound, nil
}
