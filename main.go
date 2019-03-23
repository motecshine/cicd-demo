package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// todo 通过cookies 来判断 v1 还是v2
		w.Write([]byte("v7"))
	})
	http.ListenAndServe(":8801", nil)
}
