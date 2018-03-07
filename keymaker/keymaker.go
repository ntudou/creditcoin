package main

import (
	"log"
	"net/http"

	"creditcoin/action"
)

func HttpRnd(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("t")
	if param == "" {
		result, err := action.FAAT.MakeKey()
		if err == nil {
			w.Write(result)
		}
	} else if param == "uid" {
		result, err := action.FEET.Get()
		if err == nil {
			w.Write(result)
		}
	} else {
		result, err := action.FAAT.MakeKey()
		if err == nil {
			w.Write(result)
		}
	}

}

func main() {
	http.HandleFunc("/", HttpRnd)
	err := http.ListenAndServe(":8801", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func init() {
	server_id := uint64(0)
	action.KeyMakerInit(server_id)
	action.MakeIdInit(server_id)
}
