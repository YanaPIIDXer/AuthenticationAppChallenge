package server

import (
	"fmt"
	"net/http"
	"log"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Body)
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
