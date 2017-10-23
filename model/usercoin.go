package model

import (
	"strconv"
	"creditcoin/tools"
)

type UserCoin struct {
	User_id      uint64 `json:"user_id"`
	User_coin    int64  `json:"user_coin"`
	Crc_coin     string `json:"crc_coin"`
	Limited_coin int64  `json:"limited_coin"`
	Crc_limited  string `json:"crc_limited"`
}

func NewUserCoin(user_id uint64,user_coin,limited_coin int64) (*UserCoin,error){
	uc:=&UserCoin{}
	uc.User_id=user_id
	uc.User_coin=user_coin
	uc.Limited_coin=limited_coin
	crc_coin_buf,err := tools.AESEncrypt([]byte(strconv.FormatInt(user_coin,10)),tools.AES_KEY)
	if err!=nil{
		return nil,err
	}
	uc.Crc_coin=string(crc_coin_buf)
	limited_coin_buf,err:=tools.AESEncrypt([]byte(strconv.FormatInt(limited_coin,10)),tools.AES_KEY)
	if err!=nil{
		return nil,err
	}
	uc.Crc_limited=string(limited_coin_buf)

	return uc,nil
}