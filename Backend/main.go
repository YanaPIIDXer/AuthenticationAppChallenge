package main

import(
	"fmt"
	"net/http"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
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