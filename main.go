package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// todo 通过cookies 来判断 v1 还是v2
		w.Write([]byte("v7"))
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
		}

	})
	http.ListenAndServe(":8801", nil)
}
