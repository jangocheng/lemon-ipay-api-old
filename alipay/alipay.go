package alipay

import (
	"fmt"
	"go-kit/sign"
	"io/ioutil"
	"lemon-ipay-api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kitt/mapstruct"
	alpay "github.com/relax-space/lemon-alipay-sdk"

	"github.com/labstack/echo"
	kmodel "github.com/relax-space/go-kit/model"
)

func Pay(c echo.Context) error {
	reqDto := ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.AlAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}
	if len(account.SysServiceProviderId) != 0 {
		reqDto.ExtendParams = &alpay.ExtendParams{
			SysServiceProviderId: account.SysServiceProviderId,
		}
	}
	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}

	result, err := alpay.Pay(reqDto.ReqPayDto, customDto)
	if err != nil {
		if err.Error() == alpay.MESSAGE_PAYING {
			queryDto := alpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result.OutTradeNo,
			}
			result, err = alpay.LoopQuery(&queryDto, customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
			} else {
				reverseDto := alpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result.OutTradeNo,
				}
				_, err = alpay.Reverse(&reverseDto, customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
		}
	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}

func Query(c echo.Context) error {
	reqDto := ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.AlAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	result, err := QueryCommon(account, reqDto.OutTradeNo)
	if err != nil {
		return c.JSON(http.StatusOK, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}
func Refund(c echo.Context) error {
	reqDto := ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := model.AlAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}

	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alpay.Refund(reqDto.ReqRefundDto, customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})

}
func Reverse(c echo.Context) error {
	reqDto := ReqReverseDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := model.AlAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}

	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alpay.Reverse(reqDto.ReqReverseDto, customDto, 10, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}

// func RefundQuery(c echo.Context) error {
// 	reqDto := ReqRefundQueryDto{}
// 	if err := c.Bind(&reqDto); err != nil {
// 		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
// 	}

// 	account, err := model.AlAccount{}.Get(reqDto.EId)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
// 	}
// 	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
// 		AppId:        account.AppId,
// 		AppAuthToken: account.AuthToken,
// 	}

// 	customDto := &alpay.ReqCustomerDto{
// 		PriKey: account.PriKey,
// 		PubKey: account.PubKey,
// 	}
// 	result, err := alpay.RefundQuery(reqDto.ReqRefundQueryDto, customDto)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
// 	}
// 	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
// }

func Prepay(c echo.Context) error {
	reqDto := ReqPrepayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.AlAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}
	if len(account.SysServiceProviderId) != 0 {
		reqDto.ExtendParams = &alpay.ExtendParams{
			SysServiceProviderId: account.SysServiceProviderId,
		}
	}
	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alpay.Prepay(reqDto.ReqPrepayDto, customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}

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

func Notify(c echo.Context) error {
	fmt.Printf("\n%v-%v", time.Now(), "al notify")
	sbody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Printf("\n%v-%v", time.Now(), err.Error())
		return c.String(http.StatusBadRequest, "failure")
	}
	formParam := string(sbody)
	if len(formParam) == 0 {
		fmt.Printf("\n%v-%v", time.Now(), "param is empty")
		return c.String(http.StatusBadRequest, "failure")
	}

	mapParam := base.ParseMapObjectEncode(formParam)
	var reqDto model.NotifyAlipay
	err = mapstruct.Decode(mapParam, &reqDto)
	if err != nil {
		fmt.Printf("\n%v-%v", time.Now(), err.Error())
		return c.String(http.StatusBadRequest, "failure")
	}

	//0.get account info
	nBody := reqDto.Body

	bodyMap := base.ParseMapObjectEncode(nBody)
	var eId int64
	var flag bool
	if eIdObj, ok := bodyMap["e_id"]; ok {
		if eId, err = strconv.ParseInt(eIdObj.(string), 10, 64); err == nil {
			flag = true
		}
	}
	if !flag {
		fmt.Printf("\n%v-%v", time.Now(), "e_id is not existed in param(param name:body)")
		return c.String(http.StatusBadRequest, "failure")
	}

	account, err := model.AlAccount{}.Get(eId)
	if err != nil {
		fmt.Printf("\n%v-%v", time.Now(), err.Error())
		return c.String(http.StatusBadRequest, "failure")
	}

	//1.valid sign
	signStr := reqDto.Sign
	delete(mapParam, "sign")
	delete(mapParam, "sign_type")

	if !sign.CheckSha1Sign(base.JoinMapObjectEncode(mapParam), signStr, account.PubKey) {
		fmt.Printf("\n%v-%v", time.Now(), "sign valid failure")
		return c.String(http.StatusBadRequest, "failure")
	}

	//2.valid data
	queryDto, err := QueryCommon(account, reqDto.OutTradeNo)
	if err != nil {
		fmt.Printf("\n%v-%v", time.Now(), err.Error())
		return c.String(http.StatusBadRequest, "failure")
	}
	if !(queryDto.TotalAmount == reqDto.TotalAmount) {
		fmt.Printf("\n%v-%v", time.Now(), "amount is exception")
		return c.String(http.StatusBadRequest, "failure")
	}

	//3.save notify info
	err = model.NotifyAlipay{}.InsertOne(&reqDto)
	if err != nil {
		fmt.Printf("\n%v-%v", time.Now(), err.Error())
		return c.String(http.StatusInternalServerError, "failure")
	}
	return c.String(http.StatusOK, "success")
}
