package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"

	"auth"
)

func index(w http.ResponseWriter, r *http.Request) {
	var account auth.Account
	account.Id = 1
	account.Email = "hoge@hoge"
	account.Password = "PASSWORD"

    j1, err := json.Marshal(account)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintln(w, string(j1))
}

func main() {
	fmt.Println("Start Auth Server...")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":3000", nil)
    if err != nil {
		fmt.Println("Failed...")
        log.Fatal("ListenAndServe: ", err)
    }
}
