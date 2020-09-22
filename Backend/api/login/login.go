package login_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"authapp/api/login/login_result_code"
	"authapp/api/core"
	"authapp/auth"
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
	Token auth.Token `json:"token"`
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
		result.ResultCode = login_result_code.Fatal
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
	var id = 0
	result.ResultCode, id = auth.LoginWithBasicAuth(request.Email, request.Password)
	if result.ResultCode == login_result_code.LoginSuccess {
		// トークンも取ってくる。
		token, err := auth.GetAndUpdateToken(id)
		if err == auth.TokenNotMade || err == auth.TokenIsPeriod {
			token, err = auth.MakeToken(id)
		}
		if err != nil {
			fmt.Println(err.Error())
			result.ResultCode = login_result_code.Fatal
		} else {
			result.Token = token
		}
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
