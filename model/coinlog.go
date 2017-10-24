package model

import "time"

type CoinLog struct {
	User_id          uint64    `json:"user_id"`
	Create_datetime  time.Time `json:"create_datetime"`
	Log_type		 string    `json:"log_type"`
	Pre_coin         int64     `json:"pre_coin"`
	Current_coin     int64     `json:"current_coin"`
	Pre_limited      int64     `json:"pre_limited"`
	Current_limited  int64     `json:"current_limited"`
	Log_info         string    `json:"log_info"`
}