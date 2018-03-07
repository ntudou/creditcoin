package action

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var FEET *Feet

type Feet struct {
	Serverid  uint64
	Lock      sync.Mutex
	Timestamp int64
	Autoid    int
}

func (f *Feet) Get() ([]byte, error) {
	//	buf:=bytes.Buffer{}
	var id uint64
	t := time.Now()
	n := t.Unix()
	f.Lock.Lock()
	defer f.Lock.Unlock()

	if n == f.Timestamp {
		if f.Autoid < 999999 {
			f.Autoid++
		} else {
			return nil, errors.New("Autoid is max")
		}
	} else {
		f.Timestamp = n
		f.Autoid = 0
	}

	//"18446744073709551615"
	//"9999999999999999999"
	//"200636586400"
	id = uint64(t.Year())
	id = id*1000 + uint64(t.YearDay())
	//	buf.WriteByte((byte)((yd >> 24) & 0xFF))
	//	buf.WriteByte((byte)((yd >> 16) & 0xFF))
	//	buf.WriteByte((byte)((yd >> 8) & 0xFF))
	//	buf.WriteByte((byte)((yd) & 0xFF))
	id = id*100000 + uint64(t.Hour()*3600+t.Minute()*60+t.Second())
	//	buf.WriteByte((byte)((sd >> 24) & 0xFF))
	//	buf.WriteByte((byte)((sd >> 16) & 0xFF))
	//	buf.WriteByte((byte)((sd >> 8) & 0xFF))
	//	buf.WriteByte((byte)((sd) & 0xFF))
	id = id*100 + f.Serverid
	id = id*100000 + uint64(f.Autoid)
	//	buf.WriteByte((byte)((f.Autoid >> 24) & 0xFF))
	//	buf.WriteByte((byte)((f.Autoid >> 16) & 0xFF))
	//	buf.WriteByte((byte)((f.Autoid >> 8) & 0xFF))
	//	buf.WriteByte((byte)((f.Autoid) & 0xFF))
	return []byte(strconv.FormatUint(id, 10)), nil
}

func NewFeet(id uint64) *Feet {
	f := &Feet{}
	f.Timestamp = time.Now().Unix()
	f.Autoid = 0
	f.Serverid = id
	return f
}

func MakeIdInit(id uint64) {
	FEET = NewFeet(id)
}
