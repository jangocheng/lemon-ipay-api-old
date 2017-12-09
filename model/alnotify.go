package model

import (
	"errors"
	"time"
)

type NotifyAlipay struct {
	NotifyTime string `json:"notify_time,omitempty" form:"notify_time,omitempty"`
	NotifyType string `json:"notify_type,omitempty" form:"notify_type,omitempty"`
	NotifyId   string `json:"notify_id,omitempty" form:"notify_id,omitempty"`
	SignType   string `json:"sign_type,omitempty" form:"sign_type,omitempty"`
	Sign       string `json:"sign,omitempty"  form:"sign,omitempty" xorm:"varchar(256)"`

	TradeNo    string `json:"trade_no,omitempty" form:"trade_no,omitempty"`
	AppId      string `json:"app_id,omitempty" form:"app_id,omitempty" xorm:"appid"`
	OutTradeNo string `json:"out_trade_no,omitempty" form:"out_trade_no,omitempty"`
	OutBizNo   string `json:"out_biz_no,omitempty" form:"out_biz_no,omitempty"`
	BuyerId    string `json:"buyer_id,omitempty" form:"buyer_id,omitempty"`

	BuyerLogonId string `json:"buyer_logon_id,omitempty" form:"buyer_logon_id,omitempty"`
	SellerId     string `json:"seller_id,omitempty" form:"seller_id,omitempty"`
	SellerEmail  string `json:"seller_email,omitempty" form:"seller_email,omitempty"`
	TradeStatus  string `json:"trade_status,omitempty" form:"trade_status,omitempty"`
	TotalAmount  string `json:"total_amount,omitempty" form:"total_amount,omitempty"` //float64

	ReceiptAmount  string `json:"receipt_amount,omitempty" form:"receipt_amount,omitempty"`     //float64
	InvoiceAmount  string `json:"invoice_amount,omitempty" form:"invoice_amount,omitempty"`     //float64
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" form:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string `json:"point_amount,omitempty" form:"point_amount,omitempty"`         //float64
	RefundFee      string `json:"refund_fee,omitempty" form:"refund_fee,omitempty"`             //float64

	SendBackFee string `json:"send_back_fee,omitempty" form:"send_back_fee,omitempty"` //float64
	Subject     string `json:"subject,omitempty" form:"subject,omitempty" xorm:"varchar(256)"`
	Body        string `json:"body,omitempty" form:"body,omitempty"` // xorm:"varchar(400)"
	GmtCreate   string `json:"gmt_create,omitempty" form:"gmt_create,omitempty"`
	GmtPayment  string `json:"gmt_payment,omitempty" form:"gmt_payment,omitempty"`

	GmtRefund    string    `json:"gmt_refund,omitempty" form:"gmt_refund,omitempty"`
	GmtClose     string    `json:"gmt_close,omitempty" form:"gmt_close,omitempty"`
	FundBillList string    `json:"fund_bill_list,omitempty" form:"fund_bill_list,omitempty" xorm:"varchar(512)"`
	CreatedAt    time.Time `json:"created_at" form:"created_at,omitempty" xorm:"created"`
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
