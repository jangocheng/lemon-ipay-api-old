package model

import (
	"errors"
)

type NotifyAlipay struct {
	NotifyTime string `json:"notify_time"`
	NotifyType string `json:"notify_type"`
	NotifyId   string `json:"notify_id"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign" xorm:"varchar(256)"`

	TradeNo    string `json:"trade_no"`
	AppId      string `json:"app_id" xorm:"appid"`
	OutTradeNo string `json:"out_trade_no"`
	OutBizNo   string `json:"out_biz_no"`
	BuyerId    string `json:"buyer_id"`

	BuyerLogonId string  `json:"buyer_logon_id"`
	SellerId     string  `json:"seller_id"`
	SellerEmail  string  `json:"seller_email"`
	TradeStatus  string  `json:"trade_status"`
	TotalAmount  float64 `json:"total_amount"`

	ReceiptAmount  float64 `json:"receipt_amount"`
	InvoiceAmount  float64 `json:"invoice_amount"`
	BuyerPayAmount float64 `json:"buyer_pay_amount"`
	PointAmount    float64 `json:"point_amount"`
	RefundFee      float64 `json:"refund_fee"`

	SendBackFee float64 `json:"send_back_fee"`
	Subject     string  `json:"subject" xorm:"varchar(256)"`
	Body        string  `json:"body" xorm:"varchar(400)"`
	GmtCreate   string  `json:"gmt_create"`
	GmtPayment  string  `json:"gmt_payment"`

	GmtRefund    string `json:"gmt_refund"`
	GmtClose     string `json:"gmt_close"`
	FundBillList string `json:"fund_bill_list" xorm:"varchar(512)"`
}

func (NotifyAlipay) Get(appId, outTradeNo string) (notify NotifyAlipay, err error) {
	has, err := Db.Where("appId =?", appId).And("out_trade_no=?", outTradeNo).Get(&notify)
	if err != nil {
		return
	} else if !has {
		err = errors.New("no data has found.")
		return
	}
	return
}

func (NotifyAlipay) InsertOne(notify *NotifyAlipay) (err error) {
	has, err := Db.Where("appId =?", notify.AppId).And("out_trade_no=?", notify.OutTradeNo).Get(&NotifyAlipay{})
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
