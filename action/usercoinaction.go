package action

import (
	"database/sql"
	"encoding/json"
	"sync"

	_ "github.com/lib/pq"

	"creditcoin/tools"
	"creditcoin/model"
)
type UserCoin struct {
	Lock sync.RWMutex
	PgInfo *model.PgDB
	DB *sql.DB
}

func NewUserCoin(pg *model.PgDB) (*UserCoin,error){
	var err error
	uc := & UserCoin{}
	uc.PgInfo=pg
	uc.DB, err = sql.Open("postgres", uc.PgInfo.ToString())
	if err != nil {
		return nil, err
	}
	return uc, err
}

func (uc *UserCoin) DBopen() error {
	var err error
	uc.Lock.Lock()
	defer uc.Lock.Unlock()
	if uc.DB != nil {
		uc.DB.Close()
	}
	uc.DB, err = sql.Open("postgres", uc.PgInfo.ToString())

	return err

}

