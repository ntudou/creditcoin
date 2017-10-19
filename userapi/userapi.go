package main

import (
	"net/http"
	"io/ioutil"

	"creditcoin/action"
	"log"
	"creditcoin/model"
)

func HttpRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if r.Method != "POST" {
		log.Println("request method is not post")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("request method is not post"))
		return
	}
	encode, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = action.USERAPI.Register(encode)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", HttpRegister)
	err := http.ListenAndServe(":8802", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func init() {
	var err error
	pginfo := model.NewPgDB("", "", "", "", 5000)
	action.USERAPI, err = action.NewUserApi([]byte("5J7ziRFFu7NOH00_gDSNugCj1NPNmG1h"), pginfo)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
