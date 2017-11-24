package datadb

import (
	"errors"

	wxpay "github.com/relax-space/lemon-wxpay"
)

type NotifyWechat struct {
	wxpay.NotifyDto
}

func (NotifyWechat) Get(appId, mchId, outTradeNo string) (notify *NotifyWechat, err error) {
	has, err := Db.Where("appId =?", appId).And("mch_id=?", mchId).And("out_trade_no=?", outTradeNo).Get(&notify)
	if err != nil {
		return
	} else if !has {
		err = errors.New("no data has found.")
		return
	}
	return
}

func (NotifyWechat) InsertOne(notify NotifyWechat) (err error) {
	return
}
