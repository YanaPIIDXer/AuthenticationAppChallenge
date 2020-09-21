package main

import (
	"time"
	"authapp/server"
)


func main() {
	// Timezoneを日本時間にする。
    loc, err := time.LoadLocation("Asia/Tokyo")
    if err != nil {
        loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc
	
	server.Start()
}
