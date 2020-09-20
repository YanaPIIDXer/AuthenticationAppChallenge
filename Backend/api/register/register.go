package register_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"api/core"
	"api/register/result_code"
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

	var result RegisterResult
	result.ResultCode = result_code.RegisterSuccess
	if request.Email == "" || request.Password == "" {
		result.ResultCode = result_code.EmptyParam
		result.ErrorMessage = "Has empty parameter."
		sendResult(w, result);
		return;
	}
	
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
