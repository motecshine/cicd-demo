package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world 1\n \r"))
		hostname := os.Getenv("HOSTNAME")
		w.Write([]byte("HOSTNAME:" + hostname + "\r \n"))
		cookies, err := r.Cookie("env")
		if err != nil {
			// emm 没设置cookies 那就 读dev conf吧
			config, err := ioutil.ReadFile("./configs/dev.json") // just pass the file name
			if err != nil {
				w.Write([]byte("\n \r"))
				w.Write([]byte(err.Error()))
			}
			w.Write(config)
			return
		}
		cookiesEnv := cookies.Value
		config, err := loadConfig(cookiesEnv)
		if err != nil {
			w.Write([]byte("\n \r"))
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("\n \r"))
		w.Write(config)
	})
	log.Println("run at :8801")
	http.ListenAndServe(":8801", nil)
}

func loadConfig(cookiesEnv string) ([]byte, error) {
	var path string
	switch cookiesEnv {
	case "test":
		path = "./configs/test.json"
	case "prod":
		path = "./configs/prod.json"
	case "stage":
		path = "./configs/stage.json"
	default:
		path = "./configs/dev.json"
	}
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return nil, err
	}
	return b, nil
}
