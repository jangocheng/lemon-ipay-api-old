package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	wxpay "github.com/relax-space/lemon-wxpay"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/model"
)

func WxPay(c echo.Context) error {
	reqDto := WxReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}

	result, err := wxpay.Pay(reqDto.ReqPayDto, customDto)
	if err != nil {
		if err.Error() == wxpay.MESSAGE_PAYING {
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			result, err = wxpay.LoopQuery(queryDto, customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(reverseDto, customDto, 10, 10)
				return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
			}
		} else {
			return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
		}
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func WxQuery(c echo.Context) error {
	reqDto := WxReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Query(reqDto.ReqQueryDto, customDto)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}
func WxRefund(c echo.Context) error {
	reqDto := WxReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertPathName,
		CertPathKey:  account.CertPathKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Refund(reqDto.ReqRefundDto, custDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})

}
func WxReverse(c echo.Context) error {
	reqDto := WxReqReverseDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertPathName,
		CertPathKey:  account.CertPathKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Reverse(reqDto.ReqReverseDto, custDto, 10, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func WxRefundQuery(c echo.Context) error {
	reqDto := WxReqRefundQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.RefundQuery(reqDto.ReqRefundQueryDto, customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func WxPrePay(c echo.Context) error {
	reqDto := WxReqPrePayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.PrePay(reqDto.ReqPrePayDto, customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, model.Result{Success: true, Result: result})
}

func WxNotify(c echo.Context) error {

	errResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "FAIL", ""}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}
	xmlBody := string(body)
	if len(xmlBody) == 0 {
		return c.XML(http.StatusBadRequest, errResult)
	}
	result, err := wxpay.Notify(xmlBody)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}
	return c.XML(http.StatusOK, result)
}
