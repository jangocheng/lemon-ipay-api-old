package main

import wxpay "github.com/relax-space/lemon-wxpay"

type WxReqPayDto struct {
	wxpay.ReqPayDto
	EId int64 `json:"e_id"`
}
type WxReqQueryDto struct {
	wxpay.ReqQueryDto
	EId int64 `json:"e_id"`
}
type WxReqRefundDto struct {
	wxpay.ReqRefundDto
	EId int64 `json:"e_id"`
}
type WxReqReverseDto struct {
	wxpay.ReqReverseDto
	EId int64 `json:"e_id"`
}
type WxReqRefundQueryDto struct {
	wxpay.ReqRefundQueryDto
	EId int64 `json:"e_id"`
}
type WxReqPrePayDto struct {
	wxpay.ReqPrePayDto
	EId int64 `json:"e_id"`
}
