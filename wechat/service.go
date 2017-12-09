package wechat

import (
	"encoding/xml"
	"errors"
	"fmt"
	"lemon-ipay-api/core"
	"lemon-ipay-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/httpreq"
	"github.com/relax-space/go-kit/log"
	"github.com/relax-space/go-kit/sign"
	paysdk "github.com/relax-space/lemon-wxpay-sdk"
	wxpay "github.com/relax-space/lemon-wxpay-sdk"
)

func NotifyQuery(account *model.WxAccount, outTradeNo string) (result map[string]interface{}, err error) {
	var reqDto paysdk.ReqQueryDto
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := &wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	reqDto.OutTradeNo = outTradeNo
	result, err = paysdk.Query(&reqDto, customDto)
	return
}

func NotifyValid(body, signParam, outTradeNo string, totalAmount int64, mapParam map[string]interface{}) (err error) {

	//0.get account info
	bodyMap := base.ParseMapObject(body, core.NOTIFY_BODY_SEP1, core.NOTIFY_BODY_SEP2)
	var eId int64
	var flag bool
	if eIdObj, ok := bodyMap["e_id"]; ok {
		if eId, err = strconv.ParseInt(eIdObj.(string), 10, 64); err == nil {
			flag = true
		}
	}
	if !flag {
		err = errors.New("e_id(int64) is not existed in param(param name:body) or format is not correct")
		return
	}

	account, err := model.WxAccount{}.Get(eId)
	if err != nil {
		return
	}

	//1.valid sign
	signStr := signParam
	delete(mapParam, "sign")
	fmt.Println(base.JoinMapObject(mapParam))

	if !sign.CheckMd5Sign(base.JoinMapObject(mapParam), account.Key, signStr) {
		err = errors.New("sign valid failure")
		return
	}
	//2.valid data
	queryMap, err := NotifyQuery(&account, outTradeNo)
	if err != nil {
		return
	}
	if !(queryMap["total_fee"].(string) == base.ToString(totalAmount)) {
		err = errors.New("amount is exception")
		return
	}
	//3.send data to sub_mch
	if eIdObj, ok := bodyMap["sub_notify_url"]; ok {
		go func() {
			_, err = httpreq.POST("", eIdObj.(string), mapParam, nil)
		}()
	}
	return
}

func NotifyError(c echo.Context, errMsg string) error {
	errResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "FAIL", ""}
	errResult.ReturnMsg = errMsg
	log.Error(errMsg)
	return c.XML(http.StatusBadRequest, errResult)

}
