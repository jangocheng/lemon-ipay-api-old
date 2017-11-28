package alipay

import (
	"lemon-ipay-api/model"
	"net/http"

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
	reqDto.ExtendParams = &alpay.ExtendParams{
		SysServiceProviderId: account.SysServiceProviderId,
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
	reqDto.ReqBaseDto = &alpay.ReqBaseDto{
		AppId:        account.AppId,
		AppAuthToken: account.AuthToken,
	}

	customDto := &alpay.ReqCustomerDto{
		PriKey: account.PriKey,
		PubKey: account.PubKey,
	}
	result, err := alpay.Query(reqDto.ReqQueryDto, customDto)
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
	reqDto.ExtendParams = &alpay.ExtendParams{
		SysServiceProviderId: account.SysServiceProviderId,
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

// func Notify(c echo.Context) error {

// 	errResult := struct {
// 		XMLName    xml.Name `xml:"xml"`
// 		ReturnCode string   `xml:"return_code"`
// 		ReturnMsg  string   `xml:"return_msg"`
// 	}{xml.Name{}, "FAIL", ""}

// 	body, err := ioutil.ReadAll(c.Request().Body)
// 	if err != nil {
// 		errResult.ReturnMsg = err.Error()
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}
// 	xmlBody := string(body)
// 	if len(xmlBody) == 0 {
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}
// 	notifyDto, err := SubNotify(xmlBody)
// 	if err != nil {
// 		errResult.ReturnMsg = err.Error()
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}
// 	if len(notifyDto.Attach) == 0 {
// 		errResult.ReturnMsg = "attach is required"
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}

// 	var attachObj struct {
// 		EId int64 `json:"e_id"`
// 	}
// 	err = json.Unmarshal([]byte(notifyDto.Attach), &attachObj)
// 	if err != nil {
// 		errResult.ReturnMsg = "The format of the attachment must be json and must contain e_id"
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}

// 	if attachObj.EId == 0 {
// 		errResult.ReturnMsg = "e_id is missing in attach"
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}

// 	account, err := model.AlAccount{}.Get(attachObj.EId)
// 	if err != nil {
// 		return c.JSON(http.StatusOK, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
// 	}

// 	s := structs.New(notifyDto)
// 	s.TagName = "json"
// 	mResult := s.Map()

// 	//sign
// 	signObj, ok := mResult["sign"]
// 	if !ok {
// 		errResult.ReturnMsg = "sign is missing"
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}
// 	delete(mResult, "sign")
// 	if !sign.CheckMd5Sign(base.JoinMapObject(mResult), account.Key, signObj.(string)) {
// 		errResult.ReturnMsg = "The signature is invalid"
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}

// 	err = model.NotifyWechat{}.InsertOne(&notifyDto)
// 	if err != nil {
// 		errResult.ReturnMsg = err.Error()
// 		return c.XML(http.StatusBadRequest, errResult)
// 	}

// 	successResult := struct {
// 		XMLName    xml.Name `xml:"xml"`
// 		ReturnCode string   `xml:"return_code"`
// 		ReturnMsg  string   `xml:"return_msg"`
// 	}{xml.Name{}, "SUCCESS", "OK"}
// 	return c.XML(http.StatusOK, successResult)
// }

// //sub_notify_url maybe exist in attach,
// //if sub_notify_url exist,then redirect to sub_notify_url
// func SubNotify(xmlBody string) (result model.NotifyWechat, err error) {
// 	err = xml.Unmarshal([]byte(xmlBody), &result)
// 	if err != nil {
// 		err = fmt.Errorf("%v:%v", alpay.MESSAGE_ALIPAY, err)
// 		return
// 	}

// 	if len(result.Attach) == 0 {
// 		return
// 	} else {
// 		var attachObj struct {
// 			SubNotifyUrl string `json:"sub_notify_url"`
// 		}
// 		err = json.Unmarshal([]byte(result.Attach), &attachObj)
// 		if err != nil {
// 			return
// 		}

// 		if len(attachObj.SubNotifyUrl) != 0 {
// 			go func() {
// 				_, err = httpreq.POST("", attachObj.SubNotifyUrl, result, nil)
// 			}()
// 		}
// 	}
// 	return
// }
