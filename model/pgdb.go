package model

import (
	"fmt"
)

type PgDB struct {
	DBname string
	User   string
	Pwd    string
	Host   string
	Port   uint16
}

func NewPgDB(dbname, user, pwd, host string, port uint16) *PgDB {
	p := &PgDB{}
	p.DBname = dbname
	p.User = user
	p.Pwd = pwd
	p.Host = host
	p.Port = port
	return p
}

func (p *PgDB) ToString() string {
	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=desable", p.DBname, p.User, p.Pwd, p.Host, p.Port)
}
