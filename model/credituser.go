package model

import (
	"time"
)

type CreditUser struct {
	User_id         uint64    `json:"user_id"`
	User_name       string    `json:"user_name"`
	Nike_name       string    `json:"nike_name"`
	Crc_code        string    `json:"crc_code"`
	Pwd             string    `json:"pwd"`
	Tel             string    `json:"tel"`
	User_code       string    `json:"user_code"`
	User_status     string    `json:"user_status"`
	Create_user     uint64    `json:"create_user"`
	Create_datetime time.Time `json:"create_datetime"`
}
