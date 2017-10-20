package action

import (
	"database/sql"
	"encoding/json"
	"crypto/sha1"
	"hash"
	"sync"

	_ "github.com/lib/pq"

	"creditcoin/tools"
	"creditcoin/model"
)

var USERAPI *UserApi

type UserApi struct {
	lock   sync.RWMutex
	Hash   hash.Hash
	PgInfo *model.PgDB
	DB     *sql.DB
}

func NewUserApi(key []byte, db *model.PgDB) (*UserApi, error) {
	var err error
	tools.AES_KEY = key
	ua := &UserApi{}
	ua.Hash = sha1.New()
	ua.PgInfo = db
	ua.DB, err = sql.Open("postgres", ua.PgInfo.ToString())
	if err != nil {
		return nil, err
	}
	return ua, err
}

func (ua *UserApi) DBopen() error {
	var err error
	ua.lock.Lock()
	defer ua.lock.Unlock()
	if ua.DB != nil {
		ua.DB.Close()
	}
	ua.DB, err = sql.Open("postgres", ua.PgInfo.ToString())

	return err

}

func (ua *UserApi) Register(input []byte) error {
	decode, err := tools.AESDecrypt(input, tools.AES_KEY)
	if err != nil {
		return err
	}
	user_info := &model.CreditUser{}
	err = json.Unmarshal(decode, user_info)
	if err != nil {
		return err
	}
	ua.lock.RLock()
	defer ua.lock.RUnlock()

	stmt, err := ua.DB.Prepare("INSERT INTO user_info.t_user_base (user_id,user_name,nike_name,crc_code,pwd,tel,user_code,user_status,create_user,create_datetime) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,current_timestamp)")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		ua.DBopen()
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		ua.DBopen()
		return err
	}
	return nil
}
