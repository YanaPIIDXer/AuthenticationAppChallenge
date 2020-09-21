package auth

import (
	"errors"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"

	"authapp/db"
)

// トークンを取得
func GetToken(id int) (string, error) {
	var token = ""
	err := msqldrv.Access(func(db *sql.DB) {
		var period time.Time = time.Now()
		var getRecord = func() (bool, error) {
			err := db.QueryRow("SELECT token, period FROM basic_auth where id=?", id).Scan(&token, &period)
			if err != nil {
				if err != sql.ErrNoRows { return false, err }
			}
			return (err == sql.ErrNoRows), nil
		}
		bFoundRecord, err := getRecord()
		if err != nil { return }

		var current = time.Now()
		var next = current.Add(1 * time.Minute)		// TODO:テスト用に１分にしている。テスト後は１２時間くらいにする。
		if !bFoundRecord {
			// 生成されていない
			err = errors.New("Token is not made.")
			return
		}
		if period.Before(current) {
			// まだ生きているので生存期限のみ更新
			err = db.QueryRow("UPDATE token SET period=?", next).Scan()
			if err == sql.ErrNoRows { err = nil }		// 多分こいつもRowを返さない。
			return
		}
		
		// 寿命
		err = errors.New("Token is period.")
	})

	return token, err
}

// トークン生成
func MakeToken(id int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(string(rand.Int())), bcrypt.DefaultCost)
	if err != nil { return "", err }

	var token = string(hash)
	var period = time.Now().Add(1 * time.Minute)
	err = msqldrv.Access(func(db *sql.DB) {
		err = db.QueryRow("INSERT INTO token VALUES(?, ?, ?)", id, token, period).Scan()
		if err == sql.ErrNoRows { err = nil }		// こいつはRowを返さない。
	})
	
	return token, err
}
