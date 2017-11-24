package wechat

import (
	wxpay "github.com/relax-space/lemon-wxpay"
)

type ReqPayDto struct {
	*wxpay.ReqPayDto
	EId int64 `json:"e_id"`
}
type ReqQueryDto struct {
	*wxpay.ReqQueryDto
	EId int64 `json:"e_id"`
}
type ReqRefundDto struct {
	*wxpay.ReqRefundDto
	EId int64 `json:"e_id"`
}
type ReqReverseDto struct {
	*wxpay.ReqReverseDto
	EId int64 `json:"e_id"`
}
type ReqRefundQueryDto struct {
	*wxpay.ReqRefundQueryDto
	EId int64 `json:"e_id"`
}
type ReqPrePayDto struct {
	*wxpay.ReqPrePayDto
	EId int64 `json:"e_id"`
}
