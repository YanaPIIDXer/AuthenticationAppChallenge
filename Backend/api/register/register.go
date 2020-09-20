package register_api

import (
	"fmt"
	"net/http"
	"encoding/json"

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
	var result RegisterResult
	result.ResultCode = result_code.RegisterSuccess
	
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
