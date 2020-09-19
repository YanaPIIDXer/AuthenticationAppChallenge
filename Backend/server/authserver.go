package server

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	"api"
	"api/login_result_code"
)

func login(w http.ResponseWriter, r *http.Request) {
	var result api.LoginResult
	result.ResultCode = login_result_code.LoginSuccess
    j1, err := json.Marshal(result)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(j1))
}

func Start() {
	fmt.Println("Start Auth Server...")
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":3000", nil)
    if err != nil {
		fmt.Println("Failed...")
        log.Fatal("ListenAndServe: ", err)
    }
}
