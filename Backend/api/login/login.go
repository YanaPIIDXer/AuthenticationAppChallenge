package login_api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

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
	ErrorMessage string `json:"error_message"`
}

// API実行
func API(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

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
