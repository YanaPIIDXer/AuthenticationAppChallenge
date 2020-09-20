package auth

import (
	"api/register/result_code"
)

// 基本認証登録
func RegisterBasicAuth(email string, password string) int {
	if email == "" || password == "" { return result_code.EmptyParam }
	
	return result_code.RegisterSuccess
}