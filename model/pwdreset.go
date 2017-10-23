package model

type PasswdReset struct {
	User_name string `json:"user_name"`
	Old_passwd string `json:"old_passwd"`
	New_passwd string `json:"new_passwd"`
}