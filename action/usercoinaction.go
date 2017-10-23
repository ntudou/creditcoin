package action

import (
	"database/sql"
	"encoding/json"
	"sync"

	_ "github.com/lib/pq"

	"creditcoin/tools"
	"creditcoin/model"
)

var USERCOINAPI *UserCoinApi

type UserCoinApi struct {
	Lock sync.RWMutex
	PgInfo *model.PgDB
	DB *sql.DB
}

func NewUserCoin(pg *model.PgDB) (*UserCoinApi,error){
	var err error
	uc := & UserCoinApi{}
	uc.PgInfo=pg
	uc.DB, err = sql.Open("postgres", uc.PgInfo.ToString())
	if err != nil {
		return nil, err
	}
	return uc, err
}

func (uc *UserCoinApi) DBopen() error {
	var err error
	uc.Lock.Lock()
	defer uc.Lock.Unlock()
	if uc.DB != nil {
		uc.DB.Close()
	}
	uc.DB, err = sql.Open("postgres", uc.PgInfo.ToString())

	return err
}

func (uc *UserCoinApi)GetInfo(user_id uint64)

