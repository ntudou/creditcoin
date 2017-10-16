package main

import (
	"log"
	"net/http"

	"creditcoin/action"
)

func HttpRnd(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("t")
	if param==""{
		w.Write(action.FAAT.MakeKey())
	}else if param=="uid" {
		w.Write(action.FEET.Get())
	}else {
		w.Write(action.FAAT.MakeKey())
	}

}

func main() {
	http.HandleFunc("/", HttpRnd)
	err := http.ListenAndServe(":8801", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func init(){
	server_id:=uint64(0)
	action.KeyMakerInit(server_id)
	action.MakeIdInit(server_id)
}