package main

import (
	"creditcoin/tools"
	"log"
)

func main() {
	encode:=[]byte("abcdefg")
	enbytes,err:=tools.AESEncrypt(encode,tools.AES_KEY)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	log.Println(enbytes)
	debytes,err:=tools.AESDecrypt(enbytes,tools.AES_KEY)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	log.Println(debytes)
	log.Println(string(debytes))
}
