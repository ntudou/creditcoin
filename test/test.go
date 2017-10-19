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

func init() {
	tools.AES_KEY = []byte("5J7zi9tseElj400_rIbY5RkPu54qPCsQ")
}