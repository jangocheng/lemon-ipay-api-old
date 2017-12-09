package wechat

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"lemon-ipay-api/core"
	"lemon-ipay-api/model"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/relax-space/go-kitt/random"

	"github.com/relax-space/lemon-wxmp-sdk/mpAuth"

	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/data"
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
			outTradeNo := result["out_trade_no"].(string)
			queryDto := wxpay.ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: outTradeNo,
			}
			result, err = wxpay.LoopQuery(&queryDto, &customDto, 40, 2)
			if err == nil {
				return c.JSON(http.StatusOK, kmodel.Result{Success: true, Result: result})
			} else {
				reverseDto := wxpay.ReqReverseDto{
					ReqBaseDto: reqDto.ReqBaseDto,
					OutTradeNo: outTradeNo,
				}
				if _, err = wxpay.Reverse(&reverseDto, &customDto, 10, 10); err != nil {
					return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
				} else {
					return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: "reverse sucess"}})
				}

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

func Prepay(c echo.Context) error {
	reqDto := ReqPrepayDto{}
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
	result, err := wxpay.Prepay(reqDto.ReqPrepayDto, &customDto)
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
	fmt.Printf("\n%v-%v", time.Now(), "wx notify received.")

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return NotifyError(c, err.Error())
	}
	xmlBody := string(body)
	if len(xmlBody) == 0 {
		return NotifyError(c, "xml is empty")
	}
	//1.get dto data
	var notifyDto model.NotifyWechat
	err = xml.Unmarshal([]byte(xmlBody), &notifyDto)
	if err != nil {
		return NotifyError(c, err.Error())
	}
	//1.1 get mapData
	wxData := data.New()
	err = wxData.FromXml(xmlBody)
	if err != nil {
		return NotifyError(c, err.Error())
	}
	//2.valid
	if err = NotifyValid(notifyDto.Attach, notifyDto.Sign, notifyDto.OutTradeNo, notifyDto.TotalFee, wxData.DataAttr); err != nil {
		return NotifyError(c, err.Error())

	}
	//3.save into data base
	err = model.NotifyWechat{}.InsertOne(&notifyDto)
	if err != nil {
		return NotifyError(c, err.Error())
	}

	successResult := struct {
		XMLName    xml.Name `xml:"xml"`
		ReturnCode string   `xml:"return_code"`
		ReturnMsg  string   `xml:"return_msg"`
	}{xml.Name{}, "SUCCESS", "OK"}
	return c.XML(http.StatusOK, successResult)
}

const (
	IPAY_WECHAT_PREPAY       = "IPAY_WECHAT_PREPAY"
	IPAY_WECHAT_PREPAY_INNER = "IPAY_WECHAT_PREPAY_INNER"
	IPAY_WECHAT_PREPAY_ERROR = "IPAY_WECHAT_PREPAY_ERROR"
	//IPAY_WECHAT_PREPAY_ERROR = "IPAY_WECHAT_PREPAY_ERROR"
)

/*
1.get secret by eId
2.get openId by secret
3.Prepay
*/
func PrepayEasy(c echo.Context) error {

	//appId := c.QueryParam("app_id")
	//pageUrl := c.QueryParam("page_url")
	prepay_param := c.QueryParam("prepay_param")

	reqDto := ReqPrepayEasyDto{}
	err := json.Unmarshal([]byte(prepay_param), &reqDto)
	if err != nil {
		errString := "request param format is not right"
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, errString, c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.String(http.StatusBadRequest, errString)
	}
	urlStr, err := validUrl(reqDto.PageUrl)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.String(http.StatusBadRequest, err.Error())
	}
	reqDto.PageUrl = fmt.Sprintf(urlStr, random.Uuid("")) + "?" + random.Uuid("")
	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqDto.PageUrl)
	}

	openIdUrlParam := &mpAuth.ReqDto{
		AppId:       account.AppId,
		State:       "state",
		RedirectUrl: fmt.Sprintf("%v/wx/%v", core.Env.HostUrl, "prepayopenid"),
		PageUrl:     reqDto.PageUrl,
	}
	SetCookie(IPAY_WECHAT_PREPAY_INNER, prepay_param, c)
	SetCookie(IPAY_WECHAT_PREPAY_ERROR, "", c)
	return c.Redirect(http.StatusFound, mpAuth.GetUrlForAccessToken(openIdUrlParam))
}

func PrepayOpenId(c echo.Context) error {
	code := c.QueryParam("code")
	reqUrl := c.QueryParam("reurl")
	cookie, err := c.Cookie(IPAY_WECHAT_PREPAY_INNER)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}
	param, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}
	reqDto := ReqPrepayEasyDto{}
	err = json.Unmarshal([]byte(param), &reqDto)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}
	account, err := model.WxAccount{}.Get(reqDto.EId)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}
	respDto, err := mpAuth.GetAccessTokenAndOpenId(code, account.AppId, account.Secret)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}
	reqDto.OpenId = respDto.OpenId
	//request Prepay
	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
		AppId:    account.AppId,
		SubAppId: account.SubAppId,
		MchId:    account.MchId,
		SubMchId: account.SubMchId,
	}
	customDto := wxpay.ReqCustomerDto{
		Key: account.Key,
	}
	//reqDto.Attach = "1111"
	result, err := wxpay.Prepay(reqDto.ReqPrepayDto, &customDto)
	if err != nil {
		SetCookie(IPAY_WECHAT_PREPAY_ERROR, err.Error(), c)
		SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
		SetCookie(IPAY_WECHAT_PREPAY, "", c)
		return c.Redirect(http.StatusFound, reqUrl)
	}

	prePayParam := make(map[string]interface{}, 0)
	prePayParam["package"] = "prepay_id=" + base.ToString(result["prepay_id"])
	prePayParam["timeStamp"] = base.ToString(time.Now().Unix())
	prePayParam["nonceStr"] = result["nonce_str"]
	prePayParam["signType"] = "MD5"
	prePayParam["appId"] = result["appid"]
	prePayParam["pay_sign"] = sign.MakeMd5Sign(base.JoinMapObject(prePayParam), account.Key)

	SetCookieObj(IPAY_WECHAT_PREPAY, prePayParam, c)
	SetCookie(IPAY_WECHAT_PREPAY_ERROR, "", c)

	return c.Redirect(http.StatusFound, reqUrl)

}

func validUrl(pageUrl string) (result string, err error) {
	result, err = url.QueryUnescape(pageUrl)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = errors.New("page_url miss")
		return
	}
	indexTag := strings.Index(result, "#")
	result = result[0:indexTag] + "%v?" + result[indexTag:]
	return
}
