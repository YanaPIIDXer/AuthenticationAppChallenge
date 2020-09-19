package server

import (
	"fmt"
	"net/http"
	"log"

	"api/login"
)

func Start() {
	fmt.Println("Start Auth Server...")
	var loginAPI = login_api.MakeAPIObject()
	http.HandleFunc("/login", loginAPI.OnRecv)
	err := http.ListenAndServe(":3000", nil)
    if err != nil {
		fmt.Println("Failed...")
        log.Fatal("ListenAndServe: ", err)
    }
}
