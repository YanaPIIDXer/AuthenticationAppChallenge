package api_base

import (
	"net/http"
)

// APIオブジェクト
type APIObject struct {
	IsPost bool
	Method func(http.ResponseWriter, *http.Request)
}

// APIオブジェクト生成
func MakeAPIObject(method func(http.ResponseWriter, *http.Request), bIsPost bool) APIObject {
	var instance APIObject
	instance.Method = method
	instance.IsPost = bIsPost
	return instance
}

// 接続時の処理
func (this APIObject) OnRecv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)

	if this.IsPost && r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
	}

	this.Method(w, r)
}
