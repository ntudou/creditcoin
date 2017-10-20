package model

import (
	"time"
)

type CreditUser struct {
	user_id         uint64    `json:"user_id"`
	user_name       string    `json:"user_name"`
	nike_name       string    `json:"nike_name"`
	crc_code        string    `json:"crc_code"`
	pwd             string    `json:"pwd"`
	tel             string    `json:"tel"`
	user_code       string    `json:"user_code"`
	user_status     string    `json:"user_status"`
	create_user     uint64    `json:"create_user"`
	create_datetime time.Time `json:"create_datetime"`
}
