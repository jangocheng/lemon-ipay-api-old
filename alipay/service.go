package alipay

import (
	"errors"
	"fmt"
	"lemon-ipay-api/core"
	"lemon-ipay-api/model"
	"strconv"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
	alpay "github.com/relax-space/lemon-alipay-sdk"
)

func QueryCommon(account model.AlAccount, outTradeNo string) (result *alpay.RespQueryDto, err error) {
	reqDto := ReqQueryDto{}
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}

	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	reqDto.OutTradeNo = outTradeNo
	result, err = alpay.Query(reqDto.ReqQueryDto, customDto)
	return
}

func ValidNotify(body, signParam, outTradeNo, totalAmount string, mapParam map[string]interface{}) (err error) {

	fmt.Println("body", body)
	//0.get account info
	bodyMap := base.ParseMapObjectEncode(body, "&", core.NOTIFY_BODY_SEP)
	fmt.Printf("after body:%+v", bodyMap)
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

	account, err := model.AlAccount{}.Get(eId)
	if err != nil {
		return
	}

	//1.valid sign
	signStr := signParam
	delete(mapParam, "sign")
	delete(mapParam, "sign_type")
	fmt.Printf("\test:%v", base.JoinMapObjectEncode(mapParam))
	fmt.Printf("\npubkey:%v", account.PubKey)
	fmt.Printf("\nsignStr:%v", signStr)

	if !sign.CheckSha1Sign(base.JoinMapObjectEncode(mapParam), signStr, account.PubKey) {
		err = errors.New("sign valid failure")
		return
	}

	//2.valid data
	queryDto, err := QueryCommon(account, outTradeNo)
	if err != nil {
		return
	}
	if !(queryDto.TotalAmount == totalAmount) {
		err = errors.New("amount is exception")
		return
	}
	return
}
