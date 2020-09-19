package login_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"api/login/result_code"
	"api/core"
)

// ログイン要求
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// ログイン結果
type LoginResult struct {
	ResultCode int `json:"result_code"`
	ErrorMessage string `json:"error_message"`
}

// API実行
func method(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	
	var request LoginRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		var result LoginResult
		result.ResultCode = result_code.Fatal
		result.ErrorMessage = err.Error()
		
		j1, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintln(w, string(j1))
		return
	}
	
	var result LoginResult
	result.ResultCode = result_code.LoginSuccess

	// ダミーのログイン判定。
	if request.Email != "hoge@hoge" && request.Password != "Password" {
		result.ResultCode = result_code.NotRegister;
	}
	
    j1, err := json.Marshal(result)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(j1))
}

// APIオブジェクト生成
func MakeAPIObject() api_core.APIObject {
	return api_core.MakeAPIObject(method, true)
}
