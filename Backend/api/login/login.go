package login_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"api/login/result_code"
	"api/base"
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
    j1, err := json.Marshal(result)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(j1))
}

// APIオブジェクト生成
func MakeAPIObject() api_base.APIObject {
	return api_base.MakeAPIObject(method, true)
}
