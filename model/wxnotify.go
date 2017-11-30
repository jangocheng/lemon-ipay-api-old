package model

import (
	"errors"
)

type NotifyWechat struct {
	ReturnCode string `xml:"return_code" json:"return_code" xorm:"return_code"`
	ReturnMsg  string `xml:"return_msg" json:"return_msg"  xorm:"return_msg"`
	AppId      string `xml:"appid" json:"appid" xorm:"appid"`
	MchId      string `xml:"mch_id" json:"mch_id" xorm:"mch_id"`
	SubAppId   string `xml:"sub_appid" json:"sub_appid" xorm:"sub_appid"`

	SubMchId   string `xml:"sub_mch_id" json:"sub_mch_id" xorm:"sub_mch_id"`
	DeviceInfo string `xml:"device_info" json:"device_info" xorm:"device_info"`
	NonceStr   string `xml:"nonce_str" json:"nonce_str" xorm:"nonce_str"`
	Sign       string `xml:"sign" json:"sign" xorm:"sign"`
	ResultCode string `xml:"result_code" json:"result_code" xorm:"result_code"`

	ErrCode     string `xml:"err_code" json:"err_code" xorm:"err_code"`
	ErrCodeDes  string `xml:"err_code_des" json:"err_code_des" xorm:"err_code_des"`
	OpenId      string `xml:"openid" json:"openid" xorm:"openid"`
	IsSubscribe string `xml:"is_subscribe" json:"is_subscribe" xorm:"is_subscribe"`
	SubOpenId   string `xml:"sub_openid" json:"sub_openid" xorm:"sub_openid"`

	SubIsSubscribe string `xml:"sub_is_subscribe" json:"sub_is_subscribe" xorm:"sub_is_subscribe"`
	TradeType      string `xml:"trade_type" json:"trade_type" xorm:"trade_type"`
	BankType       string `xml:"bank_type" json:"bank_type" xorm:"bank_type"`
	TotalFee       int64  `xml:"total_fee" json:"total_fee" xorm:"total_fee"` //int64
	FeeType        string `xml:"fee_type" json:"fee_type" xorm:"fee_type"`

	CashFee            int64  `xml:"cash_fee" json:"cash_fee" xorm:"cash_fee"` //int64
	CashFeeType        string `xml:"cash_fee_type" json:"cash_fee_type" xorm:"cash_fee_type"`
	SettlementTotalFee int64  `xml:"settlement_total_fee" json:"settlement_total_fee" xorm:"settlement_total_fee"` //int64
	CouponFee          int64  `xml:"coupon_fee" json:"coupon_fee" xorm:"coupon_fee"`                               //int64
	CouponCount        int64  `xml:"coupon_count" json:"coupon_count" xorm:"coupon_count"`                         //int64

	TransactionId string `xml:"transaction_id" json:"transaction_id" xorm:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no" xorm:"out_trade_no"`
	Attach        string `xml:"attach" json:"attach" xorm:"attach"`
	TimeEnd       string `xml:"time_end" json:"time_end" xorm:"time_end"`
}

func (NotifyWechat) Get(appId, mchId, outTradeNo string) (notify NotifyWechat, err error) {
	has, err := Db.Where("appid =?", appId).And("mch_id=?", mchId).And("out_trade_no=?", outTradeNo).Get(&notify)
	if err != nil {
		return
	} else if !has {
		err = errors.New("no data has found.")
		return
	}
	return
}

func (NotifyWechat) InsertOne(notify *NotifyWechat) (err error) {
	has, err := Db.Where("appid =?", notify.AppId).And("mch_id=?", notify.MchId).And("out_trade_no=?", notify.OutTradeNo).Get(&NotifyWechat{})
	if err != nil {
		return
	} else if has { //success,when data exsits
		//err = errors.New("insert failure, because data is exist")
		return
	}
	r, err := Db.InsertOne(notify)
	if err != nil {
		return
	} else if r == 0 {
		err = errors.New("no data has changed.")
		return
	}
	return
}
