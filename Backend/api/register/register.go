package register_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"authapp/api/core"
	"authapp/api/register/register_result_code"
	"authapp/auth"
)

// 登録要求
type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// 登録結果
type RegisterResult struct {
	ResultCode int `json:"result_code"`
	ErrorMessage string `json:"error_message"`
}

// API実行
func method(w http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	
	var request RegisterRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		var result RegisterResult
		result.ResultCode = register_result_code.Fatal
		result.ErrorMessage = err.Error()
		
		j1, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintln(w, string(j1))
		return
	}

	var result RegisterResult
	result.ResultCode = register_result_code.RegisterSuccess
	var resultCode = auth.RegisterBasicAuth(request.Email, request.Password)
	if resultCode != register_result_code.RegisterSuccess {
		var message = "Fatal Error."
		switch(resultCode) {
			case register_result_code.EmptyParam:
				message = "Has empty parameter."
				break
			case register_result_code.UsedEmail:
				message = "This email is already used."
				break
		}
		result.ErrorMessage = message
	}
	
	result.ResultCode = resultCode
	sendResult(w, result);
}

// 結果送信
func sendResult(w http.ResponseWriter, result RegisterResult) {
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
