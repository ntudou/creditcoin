package action

import (
	"time"
	"bytes"
	"strconv"
	"math/rand"
	"sync"
)

var KM map[uint64]string
var FAAT *Faat

type Faat struct{
	Serverid uint64
	Lock sync.Mutex
	Timestamp int64
	Autoid uint64
}

func NewFaat(id uint64) *Faat {
	f:=&Faat{}
	f.Timestamp=time.Now().UnixNano()
	f.Autoid=0
	f.Serverid=id
	return f
}

func (f *Faat) MakeKey() []byte {
	t := time.Now()
	n:=t.UnixNano()
	lst := bytes.Buffer{}
	var ys []uint64
	f.Lock.Lock()
	defer f.Lock.Unlock()

	if n==f.Timestamp{
		f.Autoid++
	}else{
		f.Timestamp=n
		f.Autoid=0
	}
	//9999999999999999999
	//20060102150405
	ymd, _ := strconv.ParseUint(t.Format("20060102150405"), 10, 64)
	nsd := uint64(t.Nanosecond())
	aid:=f.Autoid
	for{
		ys = append(ys, aid%uint64(62))
		aid=aid / uint64(62)
		if aid==0{
			break
		}
	}
	ys = append(ys, f.Serverid%uint64(62))
	//log.Println(nsd)
	for {
		ys = append(ys, nsd%uint64(62))
		nsd = nsd / uint64(62)
		if nsd==0{
			break
		}
	}
	for {
		ys = append(ys, ymd%uint64(62))
		ymd = ymd / uint64(62)
		if ymd==0{
			break
		}
	}
	l := len(ys)
	//log.Println(ys)
	for ; l > 0; l-- {
		lst.WriteString(KM[ys[l-1]])
	}

	lst.WriteString("_")

	for i := 0; i < (31-l); i++ {
		lst.WriteString(KM[uint64(rand.Intn(62))])
	}

	return lst.Bytes()
}

func KeyMakerInit(id uint64){
	FAAT = NewFaat(id)
	rand.Seed(time.Now().UnixNano())
	KM = map[uint64]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
		16: "g",
		17: "h",
		18: "i",
		19: "j",
		20: "k",
		21: "l",
		22: "m",
		23: "n",
		24: "o",
		25: "p",
		26: "q",
		27: "r",
		28: "s",
		29: "t",
		30: "u",
		31: "v",
		32: "w",
		33: "x",
		34: "y",
		35: "z",
		36: "A",
		37: "B",
		38: "C",
		39: "D",
		40: "E",
		41: "F",
		42: "G",
		43: "H",
		44: "I",
		45: "J",
		46: "K",
		47: "L",
		48: "M",
		49: "N",
		50: "O",
		51: "P",
		52: "Q",
		53: "R",
		54: "S",
		55: "T",
		56: "U",
		57: "V",
		58: "W",
		59: "X",
		60: "Y",
		61: "Z",
	}
}