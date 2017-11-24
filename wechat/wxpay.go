package wechat

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"lemon-epay/datadb"
	"net/http"
	"time"

	"github.com/relax-space/go-kit/sign"

	"github.com/relax-space/go-kit/base"

	wxpay "github.com/relax-space/lemon-wxpay"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/model"
)

func Pay(c echo.Context) error {
	reqDto := ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := datadb.Account{}.Get(reqDto.EId)
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

func Query(c echo.Context) error {
	reqDto := ReqQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := datadb.Account{}.Get(reqDto.EId)
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
func Refund(c echo.Context) error {
	reqDto := ReqRefundDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := datadb.Account{}.Get(reqDto.EId)
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
func Reverse(c echo.Context) error {
	reqDto := ReqReverseDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}
	account, err := datadb.Account{}.Get(reqDto.EId)
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

func RefundQuery(c echo.Context) error {
	reqDto := ReqRefundQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := datadb.Account{}.Get(reqDto.EId)
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

func PrePay(c echo.Context) error {
	reqDto := ReqPrePayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := datadb.Account{}.Get(reqDto.EId)
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

	prePayParam := make(map[string]interface{}, 0)
	prePayParam["package"] = "prepay_id=" + base.ToString(result["prepay_id"])
	prePayParam["timeStamp"] = base.ToString(time.Now().Unix())
	prePayParam["nonceStr"] = result["nonce_str"]
	prePayParam["signType"] = "MD5"
	prePayParam["appId"] = result["appid"]
	prePayParam["pay_sign"] = sign.MakeMd5Sign(base.JoinMapObject(prePayParam), account.Key)

	return c.JSON(http.StatusOK, model.Result{Success: true, Result: prePayParam})
}

func Notify(c echo.Context) error {

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
	_, mResult, err := wxpay.Notify(xmlBody)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}

	attach, ok := mResult["attach"]
	if !ok {
		errResult.ReturnMsg = "attach is required"
		return c.XML(http.StatusBadRequest, errResult)
	}
	var attachObj struct {
		EId int64 `json:"e_id"`
	}
	err = json.Unmarshal([]byte(attach.(string)), &attachObj)
	if err != nil {
		errResult.ReturnMsg = "The format of the attachment must be json and must contain e_id"
		return c.XML(http.StatusBadRequest, errResult)
	}

	account, err := datadb.Account{}.Get(attachObj.EId)
	if err != nil {
		return c.JSON(http.StatusOK, model.Result{Success: false, Error: model.Error{Code: 10004, Message: err.Error()}})
	}

	//sign
	signObj, ok := mResult["sign"]
	if !ok {
		errResult.ReturnMsg = "sign is missing"
		return c.XML(http.StatusBadRequest, errResult)
	}
	delete(mResult, "sign")
	if !sign.CheckSign(base.JoinMapObject(mResult), account.Key, signObj.(string)) {
		errResult.ReturnMsg = "The signature is invalid"
		return c.XML(http.StatusBadRequest, errResult)
	}

	successResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "SUCCESS", "OK"}
	return c.XML(http.StatusOK, successResult)
}
