package login_api

import (
	"fmt"
	"net/http"
	"encoding/json"

	"api/login/result_code"
)

// ログイン要求
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// ログイン結果
type LoginResult struct {
	ResultCode int `json:"result_code"`
}

// API実行
func API(w http.ResponseWriter, r *http.Request) {
	var result LoginResult
	result.ResultCode = result_code.LoginSuccess
    j1, err := json.Marshal(result)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(j1))
}
