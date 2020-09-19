package api

import {
	"net/http"
}

type IAPI interface {
	API(w http.ResponseWriter, r *http.Request)
}

type APIBase struct {
	var bIsPost bool = false
	var API IAPI = nil
}

// 接続時の処理
func (this APIBase) OnRecv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	if this.bIsPost && r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}

	this.API(w, r)
}
