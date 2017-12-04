package wechat

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
)

func SetCookie(key, value string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = key
	value = url.QueryEscape(value)
	cookie.Value = value
	cookie.Domain = "p2shop.cn"
	cookie.Path = "/"
	//cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
}

func SetCookieObj(key string, value interface{}, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = key
	b, _ := json.Marshal(value)
	cookie.Value = url.QueryEscape(string(b))
	cookie.Domain = "p2shop.cn"
	cookie.Path = "/"
	//cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
}

// func validUrl(pageUrl string) (result string, err error) {
// 	result, err = url.QueryUnescape(pageUrl)
// 	if err != nil {
// 		return
// 	}
// 	if len(result) == 0 {
// 		err = errors.New("page_url miss")
// 		return
// 	}
// 	indexTag := strings.Index(result, "#")
// 	result = result[0:indexTag] + "%v?" + result[indexTag:]
// 	return
// }

// func setErrorCookie(errMsg string, c echo.Context) {
// 	SetCookie(IPAY_WECHAT_PREPAY_ERROR, errMsg, c)
// 	SetCookie(IPAY_WECHAT_PREPAY_INNER, "", c)
// 	SetCookie(IPAY_WECHAT_PREPAY, "", c)
// }

// func GetPrepayParam(reqDto *wxpay.ReqPrepayDto, account *model.WxAccount) (prePayParam map[string]interface{}, err error) {
// 	//request Prepay
// 	reqDto.ReqBaseDto = &wxpay.ReqBaseDto{
// 		AppId:    account.AppId,
// 		SubAppId: account.SubAppId,
// 		MchId:    account.MchId,
// 		SubMchId: account.SubMchId,
// 	}
// 	customDto := wxpay.ReqCustomerDto{
// 		Key: account.Key,
// 	}
// 	reqDto.Attach = "1111"
// 	result, err := wxpay.Prepay(reqDto, &customDto)
// 	if err != nil {
// 		return
// 	}

// 	prePayParam["package"] = "prepay_id=" + base.ToString(result["prepay_id"])
// 	prePayParam["timeStamp"] = base.ToString(time.Now().Unix())
// 	prePayParam["nonceStr"] = result["nonce_str"]
// 	prePayParam["signType"] = "MD5"
// 	prePayParam["appId"] = result["appid"]
// 	prePayParam["pay_sign"] = sign.MakeMd5Sign(base.JoinMapObject(prePayParam), account.Key)
// 	return
// }
