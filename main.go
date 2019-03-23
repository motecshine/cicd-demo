package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("v7 "))
	})
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println(`{"v7Timestamp":"2018-11-07 14:12:57","Level":"DBUG","File":"heartbeat:48","ID":"","Action":"","Message":"更新心跳map[]"}`)
			}
		}
	}()
	http.ListenAndServe(":8801", nil)
}
