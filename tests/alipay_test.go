package tests

import (
	"encoding/json"
	"fmt"
	"kit/test"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"lemon-ipay-api/alipay"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/model"
)

func Test_AliPay(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"auth_code":"283209675485586567",
		"subject":"xiaomiao test apilay",
		"total_amount":0.01
	}`
	req, err := http.NewRequest(echo.POST, "/v3/al/pay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, alipay.Pay(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_AliRefund(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"1117112912739763007486053235",
		"refund_amount":0.01
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/refund", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, alipay.Refund(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_AliQuery(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"1117112912739763007486053235"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/query", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, alipay.Query(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_AliReverse(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"out_trade_no":"1117112912739763007486053235"
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/reverse", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, alipay.Reverse(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}

func Test_AliPrepay(t *testing.T) {
	bodyStr := `
	{
		"e_id":10001,
		"subject":"xiaomiao test ali",
		"total_amount":0.01
	}`
	req, err := http.NewRequest(echo.POST, "/v3/wx/prepay", strings.NewReader(bodyStr))
	test.Ok(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	test.Ok(t, alipay.Prepay(c))
	v := model.Result{}
	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	fmt.Printf("%+v", v)
	test.Equals(t, http.StatusOK, rec.Code)

}
