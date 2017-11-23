package main

import "errors"

type WxAccount struct {
	AppId        string
	SubAppId     string
	Key          string
	MchId        string
	SubMchId     string
	CertPathName string
	CertPathKey  string
	RootCa       string
}

func (WxAccount) TableName() string {
	return "wechat"
}

func (WxAccount) Get(eId int64) (account WxAccount, err error) {
	has, err := db.Where("e_id =?", eId).Get(&account)
	if err != nil {
		return
	} else if !has {
		err = errors.New("no data has found.")
		return
	}
	return
}
