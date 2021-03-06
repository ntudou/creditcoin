package action

import (
	"bytes"
	"crypto/sha512"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"errors"
	"hash"
	"io/ioutil"
	"net/http"
	"sync"

	_ "github.com/lib/pq"

	"creditcoin/model"
	"creditcoin/tools"
)

var USERAPI *UserApi

type UserApi struct {
	UserID_URL   string
	UserCode_URL string
	Lock         sync.RWMutex
	Hash         hash.Hash
	PgInfo       *model.PgDB
	DB           *sql.DB
}

func NewUserApi(key []byte, userid_url, usercode_url string, db *model.PgDB) (*UserApi, error) {
	var err error
	tools.AES_KEY = key
	ua := &UserApi{}
	ua.UserID_URL = userid_url
	ua.UserCode_URL = usercode_url
	ua.Hash = sha512.New()
	ua.PgInfo = db
	ua.DB, err = sql.Open("postgres", ua.PgInfo.ToString())
	if err != nil {
		return nil, err
	}
	return ua, err
}

func (ua *UserApi) DBopen() error {
	var err error
	ua.Lock.Lock()
	defer ua.Lock.Unlock()
	if ua.DB != nil {
		ua.DB.Close()
	}
	ua.DB, err = sql.Open("postgres", ua.PgInfo.ToString())

	return err

}

func (ua *UserApi) UserNameSearch(input []byte) (bool, error) {
	decode, err := tools.AESDecrypt(input, tools.AES_KEY)
	if err != nil {
		return false, err
	}
	user_name := string(decode)
	rows, err := ua.DB.Query("SELECT count(1) FROM user_info.t_user_base WHERE user_name=$1", user_name)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		ua.DBopen()
		return false, err
	}
	if rows.Next() {
		var user_count uint64
		err = rows.Scan(&user_count)
		if err != nil {
			return false, err
		}
		if user_count < 1 {
			return true, nil
		} else {
			return false, errors.New("user name is exist")
		}
	} else {
		return false, errors.New("result is empty")
	}
}

func (ua *UserApi) InputPwd(pr *model.PasswdReset) (bool, error) {
	resp, err := http.Get(ua.UserCode_URL)
	defer resp.Body.Close()
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	crc_encode, err := tools.AESEncrypt(body, tools.AES_KEY)
	if err != nil {
		return false, err
	}
	str_crc_encode := string(crc_encode)

	pwd_buf := &bytes.Buffer{}
	pwd_buf.WriteString(pr.New_passwd)
	pwd_buf.Write(body)
	pwd_sha := ua.Hash.Sum(pwd_buf.Bytes())
	pwd_str := string(pwd_sha)

	_, err = ua.DB.Exec("UPDATE user_info.t_user_base SET crc_code=$1,pwd=$2 WHERE user_name=$3", str_crc_encode, pwd_str, pr.User_name)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (ua *UserApi) PwdReset(input []byte) (bool, error) {
	decode, err := tools.AESDecrypt(input, tools.AES_KEY)
	if err != nil {
		return false, err
	}
	pr := model.PasswdReset{}
	err = json.Unmarshal(decode, pr)
	if err != nil {
		return false, err
	}
	rows, err := ua.DB.Query("SELECT crc_code,pwd FROM user_info.t_user_base WHERE user_name=$1", pr.User_name)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		ua.DBopen()
		return false, err
	}

	if rows.Next() {
		var crc_code string
		var pwd string
		err = rows.Scan(&crc_code, &pwd)
		if err != nil {
			return false, err
		}
		crc_decode, err := tools.AESDecrypt([]byte(crc_code), tools.AES_KEY)
		if err != nil {
			return false, err
		}

		pwd_buf := &bytes.Buffer{}
		pwd_buf.WriteString(pr.Old_passwd)
		pwd_buf.Write(crc_decode)
		pwd_sha := ua.Hash.Sum(pwd_buf.Bytes())

		if string(pwd_sha) == pwd {
			return ua.InputPwd(&pr)
		} else {
			return false, errors.New("Incorrect user name/password")
		}

	} else {
		return false, errors.New("Incorrect user name/password")
	}
}

func (ua *UserApi) Login(input []byte) (bool, error) {
	decode, err := tools.AESDecrypt(input, tools.AES_KEY)
	if err != nil {
		return false, err
	}
	user_login := &model.Userlogin{}
	err = json.Unmarshal(decode, user_login)
	if err != nil {
		return false, err
	}

	rows, err := ua.DB.Query("SELECT crc_code,pwd FROM user_info.t_user_base WHERE user_name=$1", user_login.User_name)

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		ua.DBopen()
		return false, err
	}

	if rows.Next() {
		var crc_code string
		var pwd string
		err = rows.Scan(&crc_code, &pwd)
		if err != nil {
			return false, err
		}
		crc_decode, err := tools.AESDecrypt([]byte(crc_code), tools.AES_KEY)
		if err != nil {
			return false, err
		}

		pwd_buf := &bytes.Buffer{}
		pwd_buf.WriteString(user_login.Pwd)
		pwd_buf.Write(crc_decode)
		pwd_sha := ua.Hash.Sum(pwd_buf.Bytes())

		if string(pwd_sha) == pwd {
			return true, nil
		} else {
			return false, errors.New("Incorrect user name/password")
		}

	} else {
		return false, errors.New("Incorrect user name/password")
	}

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

	resp, err := http.Get(ua.UserCode_URL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	user_crc_code, err := tools.AESEncrypt(body, tools.AES_KEY)
	if err != nil {
		return err
	}
	user_info.Crc_code = string(user_crc_code)

	user_id_resp, err := http.Get(ua.UserID_URL)
	defer user_id_resp.Body.Close()
	if err != nil {
		return err
	}

	user_id_body, err := ioutil.ReadAll(user_id_resp.Body)
	if err != nil {
		return err
	}

	user_info.User_id = binary.BigEndian.Uint64(user_id_body)

	pwd_buf := &bytes.Buffer{}
	pwd_buf.WriteString(user_info.Pwd)
	pwd_buf.Write(body)
	pwd_sha := ua.Hash.Sum(pwd_buf.Bytes())
	pwd_str := string(pwd_sha)

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

	_, err = stmt.Exec(user_info.User_id, user_info.User_name, user_info.Nike_name, user_info.Crc_code, pwd_str, user_info.Tel, user_info.User_code, user_info.User_status, user_info.Create_user)
	if err != nil {
		ua.DBopen()
		return err
	}
	return nil
}
