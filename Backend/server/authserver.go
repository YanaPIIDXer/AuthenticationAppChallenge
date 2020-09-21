package server

import (
	"fmt"
	"net/http"
	"log"
	"math/rand"
	"time"

	"authapp/api/login"
	"authapp/api/register"
)

func Start() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Start Auth Server...")
	var loginAPI = login_api.MakeAPIObject()
	http.HandleFunc("/login", loginAPI.OnRecv)

	var registerAPI = register_api.MakeAPIObject()
	http.HandleFunc("/register", registerAPI.OnRecv)
	err := http.ListenAndServe(":3000", nil)
    if err != nil {
		fmt.Println("Failed...")
        log.Fatal("ListenAndServe: ", err)
    }
}
