package server

import (
	"fmt"
	"net/http"
	"log"

	"api/login"
)

func Start() {
	fmt.Println("Start Auth Server...")
	http.HandleFunc("/login", login_api.API)
	err := http.ListenAndServe(":3000", nil)
    if err != nil {
		fmt.Println("Failed...")
        log.Fatal("ListenAndServe: ", err)
    }
}
