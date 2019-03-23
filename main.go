package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// read cookie
		var cookie, err = r.Cookie("env")
		if err == nil {
			var cookievalue = cookie.Value
			io.WriteString(w, "<b>get cookie value is "+cookievalue+"</b>\n")
			switch cookievalue {
			case "prod":
				w.Write([]byte("prod"))
			case "test":
				w.Write([]byte("test"))
			case "stage":
				w.Write([]byte("stage"))
			default:
				w.Write([]byte("cookievalue"))
			}
		} else {
			w.Write([]byte(err.Error()))
		}

	})
	http.ListenAndServe(":8801", nil)
}
