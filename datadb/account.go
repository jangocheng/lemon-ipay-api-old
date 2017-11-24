package datadb

import "errors"

type Account struct {
	AppId        string
	SubAppId     string
	Key          string
	MchId        string
	SubMchId     string
	CertPathName string
	CertPathKey  string
	RootCa       string
}

func (Account) TableName() string {
	return "wechat"
}

func (Account) Get(eId int64) (account Account, err error) {
	has, err := Db.Where("e_id =?", eId).Get(&account)
	if err != nil {
		return
	} else if !has {
		err = errors.New("no data has found.")
		return
	}
	return
}
