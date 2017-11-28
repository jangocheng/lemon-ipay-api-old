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

func SetCookieObj(key, value string, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = key
	b, _ := json.Marshal(value)
	value = url.QueryEscape(string(b))
	cookie.Value = value
	cookie.Domain = "p2shop.cn"
	cookie.Path = "/"
	//cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)
}
