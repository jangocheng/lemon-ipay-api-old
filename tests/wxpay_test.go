package tests

import (
	"encoding/json"
	"fmt"
	"kit/test"
	"lemon-ipay-api/wechat"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/model"
)

func Test_WxPay(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"auth_code":"135298324463700425",
		"body":"xiaoxinmiao test",
		"total_fee":1
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/pay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.Pay(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxRefund(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"147688874645492354650",
		"refund_fee":1
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/refund", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.Refund(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxReverse(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"143420620288156126697"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/reverse", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.Reverse(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxQuery(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"14201711085205823413229775520"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/query", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.Query(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxRefundQuery(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"144650782494807835413"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/refundquery", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.RefundQuery(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxPrePay(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"body":"xiaomiao test",
		"total_fee":1,
		"trade_type":"JSAPI",
		"notify_url":"http://xiao.xinmiao.com",
		"openid":"os2u9uPKLkCKL08FwCM6hQAQ_LtI"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/prepay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.PrePay(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_WxNotify(t *testing.T) {
	xmlBody := `<xml>
	<appid><![CDATA[wx2421b1c4370ec43b]]></appid>
	<attach><![CDATA[{"sub_notify_url":"https://baidu.com","e_id":10001}]]></attach>
	<bank_type><![CDATA[CFT]]></bank_type>
	<fee_type><![CDATA[CNY]]></fee_type>
	<is_subscribe><![CDATA[Y]]></is_subscribe>
	<mch_id><![CDATA[10000100]]></mch_id>
	<nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
	<openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
	<out_trade_no><![CDATA[1409811653]]></out_trade_no>
	<result_code><![CDATA[SUCCESS]]></result_code>
	<return_code><![CDATA[SUCCESS]]></return_code>
	<sign><![CDATA[CDF756E879A802281624B691BEAEAEF4]]></sign>
	<sub_mch_id><![CDATA[10000100]]></sub_mch_id>
	<time_end><![CDATA[20140903131540]]></time_end>
	<total_fee>1</total_fee>
	<trade_type><![CDATA[JSAPI]]></trade_type>
	<transaction_id><![CDATA[B2AE05C99B9C81A640472406AA3C2710]]></transaction_id>
 </xml>`
	req, err := http.NewRequest(echo.POST, "/v3/wx/notify", strings.NewReader(xmlBody))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, wechat.Notify(c))
	fmt.Printf("%+v", string(rec.Body.Bytes()))
	test.Equals(t, http.StatusOK, rec.Code)

}
