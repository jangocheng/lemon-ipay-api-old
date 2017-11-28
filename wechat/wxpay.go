package wechat

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"lemon-ipay-api/core"
	"lemon-ipay-api/model"
	"net/http"
	"net/url"
	"time"

	"github.com/relax-space/lemon-wxmp-sdk/mpAuth"

	"github.com/relax-space/go-kit/httpreq"

	"github.com/fatih/structs"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"

	wxpay "github.com/relax-space/lemon-wxpay-sdk"

	"github.com/labstack/echo"
	kmodel "github.com/relax-space/go-kit/model"
)

func Pay(c echo.Context) error {
	reqDto := ReqPayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}

	result, err := wxpay.Pay(reqDto.ReqPayDto, &customDto)
	if err != nil {
		if err.Error() == wxpay.MESSAGE_PAYING {
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result["out_trade_no"].(string),
			}
			result, err = wxpay.LoopQuery(&queryDto, &customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: result["out_trade_no"].(string),
				}
				_, err = wxpay.Reverse(&reverseDto, &customDto, 10, 10)
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

	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.Query(reqDto.ReqQueryDto, &customDto)
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
	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertName,
		CertPathKey:  account.CertKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Refund(reqDto.ReqRefundDto, &custDto)
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
	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	custDto := wxpay.ReqCustomerDto{
		Key:          account.Key,
		CertPathName: account.CertName,
		CertPathKey:  account.CertKey,
		RootCa:       account.RootCa,
	}
	result, err := wxpay.Reverse(reqDto.ReqReverseDto, &custDto, 10, 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})

	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}

func RefundQuery(c echo.Context) error {
	reqDto := ReqRefundQueryDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.RefundQuery(reqDto.ReqRefundQueryDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
}

func PrePay(c echo.Context) error {
	reqDto := ReqPrePayDto{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.PrePay(reqDto.ReqPrePayDto, &customDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	prePayParam := make(map[string]interface{}, 0)
	prePayParam["package"] = "prepay_id=" + base.ToString(result["prepay_id"])
	prePayParam["timeStamp"] = base.ToString(time.Now().Unix())
	prePayParam["nonceStr"] = result["nonce_str"]
	prePayParam["signType"] = "MD5"
	prePayParam["appId"] = result["appid"]
	prePayParam["pay_sign"] = sign.MakeMd5Sign(base.JoinMapObject(prePayParam), account.Key)

	return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: prePayParam})
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
	notifyDto, err := SubNotify(xmlBody)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}
	if len(notifyDto.Attach) == 0 {
		errResult.ReturnMsg = "attach is required"
		return c.XML(http.StatusBadRequest, errResult)
	}

	var attachObj struct {
		EId int64 `json:"e_id"`
	}
	err = json.Unmarshal([]byte(notifyDto.Attach), &attachObj)
	if err != nil {
		errResult.ReturnMsg = "The format of the attachment must be json and must contain e_id"
		return c.XML(http.StatusBadRequest, errResult)
	}

	if attachObj.EId == 0 {
		errResult.ReturnMsg = "e_id is missing in attach"
		return c.XML(http.StatusBadRequest, errResult)
	}

	account, err := model.WxAccount{}.Get(attachObj.EId)
	if err != nil {
		return c.JSON(http.StatusOK, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	s := structs.New(notifyDto)
	s.TagName = "json"
	mResult := s.Map()

	//sign
	signObj, ok := mResult["sign"]
	if !ok {
		errResult.ReturnMsg = "sign is missing"
		return c.XML(http.StatusBadRequest, errResult)
	}
	delete(mResult, "sign")
	if !sign.CheckMd5Sign(base.JoinMapObject(mResult), account.Key, signObj.(string)) {
		errResult.ReturnMsg = "The signature is invalid"
		return c.XML(http.StatusBadRequest, errResult)
	}

	err = model.NotifyWechat{}.InsertOne(&notifyDto)
	if err != nil {
		errResult.ReturnMsg = err.Error()
		return c.XML(http.StatusBadRequest, errResult)
	}

	successResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "SUCCESS", "OK"}
	return c.XML(http.StatusOK, successResult)
}

//sub_notify_url maybe exist in attach,
//if sub_notify_url exist,then redirect to sub_notify_url
func SubNotify(xmlBody string) (result model.NotifyWechat, err error) {
	err = xml.Unmarshal([]byte(xmlBody), &result)
	if err != nil {
		err = fmt.Errorf("%v:%v", wxpay.MESSAGE_WECHAT, err)
		return
	}

	if len(result.Attach) == 0 {
		return
	} else {
		var attachObj struct {
			SubNotifyUrl string `json:"sub_notify_url"`
		}
		err = json.Unmarshal([]byte(result.Attach), &attachObj)
		if err != nil {
			return
		}

		if len(attachObj.SubNotifyUrl) != 0 {
			go func() {
				_, err = httpreq.POST("", attachObj.SubNotifyUrl, result, nil)
			}()
		}
	}
	return
}

const (
	IPAY_WECHAT_PREPAY       = "IPAY_WECHAT_PREPAY"
	IPAY_WECHAT_PREPAY_ERROR = "IPAY_WECHAT_PREPAY_ERROR"
	//IPAY_WECHAT_PREPAY_ERROR = "IPAY_WECHAT_PREPAY_ERROR"
)

/*
1.get secret by eId
2.get openId by secret
3.prepay
*/
func PrePayEasy(c echo.Context) error {

	appId := c.QueryParam("app_id")
	pageUrl := c.QueryParam("page_url")
	prepay_param := c.QueryParam("prepay_param")

	openIdUrlParam := &mpAuth.ReqDto{
		AppId:       appId,
		State:       "state",
		RedirectUrl: fmt.Sprintf("%v/wx/%v", core.Env.HostUrl, "prepayopenid"),
		PageUrl:     pageUrl,
	}
	reqUrl := mpAuth.GetUrlForAccessToken(openIdUrlParam)

	SetCookie(IPAY_WECHAT_PREPAY, prepay_param, c)
	return c.Redirect(http.StatusFound, reqUrl)
}

func PrepayOpenId(c echo.Context) error {
	code := c.QueryParam("code")
	reqUrl := c.QueryParam("reurl")
	//	param := c.Param("param")

	cookie, err := c.Cookie(IPAY_WECHAT_PREPAY)
	fmt.Printf("11111%+v,%v", cookie, err)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}
	param, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}
	reqDto := ReqPrePayDto{}
	err = json.Unmarshal([]byte(param), &reqDto)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}
	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}
	respDto, err := mpAuth.GetAccessTokenAndOpenId(code, account.AppId, account.Secret)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}
	reqDto.OpenId = respDto.OpenId

	//request prepay
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	result, err := wxpay.PrePay(reqDto.ReqPrePayDto, &customDto)
	if err != nil {
		q := make(url.Values)
		q.Set(IPAY_WECHAT_PREPAY_ERROR, err.Error())
		return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())
	}

	prePayParam := make(map[string]interface{}, 0)
	prePayParam["package"] = "prepay_id=" + base.ToString(result["prepay_id"])
	prePayParam["timeStamp"] = base.ToString(time.Now().Unix())
	prePayParam["nonceStr"] = result["nonce_str"]
	prePayParam["signType"] = "MD5"
	prePayParam["appId"] = result["appid"]
	prePayParam["pay_sign"] = sign.MakeMd5Sign(base.JoinMapObject(prePayParam), account.Key)

	b, _ := json.Marshal(prePayParam)
	prepayParamStr := url.QueryEscape(string(b))
	q := make(url.Values)
	q.Set(IPAY_WECHAT_PREPAY, prepayParamStr)
	return c.Redirect(http.StatusFound, reqUrl+"?"+q.Encode())

}
