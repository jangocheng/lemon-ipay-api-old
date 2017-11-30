package alipay

import (
	"fmt"
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

func Notify(c echo.Context) error {
	reqDto := model.NotifyAlipay{}
	if err := c.Bind(&reqDto); err != nil {
		return c.JSON(http.StatusBadRequest, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}
	fmt.Printf("notify:%+v", reqDto)
	// account, err := model.AlAccount{}.GetByAppId(reqDto.AppId)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	// }

	// customDto := &alpay.ReqCustomerDto{
	// 	PriKey: account.PriKey,
	// 	PubKey: account.PubKey,
	// }
	// err = alpay.CheckNotifySign(reqDto.ReqNotifyDto, customDto)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	// }
	err := model.NotifyAlipay{}.InsertOne(&reqDto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, kmodel.Result{Success: false, Error: kmodel.Error{Code: 10004, Message: err.Error()}})
	}

	return c.String(http.StatusOK, "success")
}
