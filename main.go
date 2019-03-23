package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname := os.Getenv("HOSTNAME")
		w.Write([]byte(hostname))
	})
	http.ListenAndServe(":8801", nil)
}
