package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appEnv      = flag.String("APP_ENV", os.Getenv("APP_ENV"), "APP_ENV")
	connEnv     = flag.String("CONN_ENV", os.Getenv("CONN_ENV"), "CONN_ENV")
	bmappingUrl = flag.String("BMAPPING_URL", os.Getenv("BMAPPING_URL"), "BMAPPING_URL")
	envParam    EnvParamDto
	db          *xorm.Engine
)

func init() {
	flag.Parse()
	envParam = EnvParamDto{
		AppEnv:      *appEnv,
		ConnEnv:     *connEnv,
		BmappingUrl: *bmappingUrl,
	}
	//initTest()
	db = InitDB("mysql", envParam.ConnEnv)
	db.Sync(new(WxAccount))
}

func main() {

	e := echo.New()
	e.Use(middleware.CORS())
	RegisterApi(e)
	e.Start(":5000")
}

func RegisterApi(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "lemon epay")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//util.Lang = c.Request().Header["Accept-Language"]
			return next(c)
		}
	}

	v3 := e.Group("/v3", track)
	v3.POST("/pay", Pay)
	v3.POST("/query", Query)
	v3.POST("/reverse", Reverse)
	v3.POST("/refund", Refund)
	v3.POST("/prepay", PrePay)

	v3.GET("/record/:Id", GetRecord)

	wx := v3.Group("/wx")
	wx.POST("/pay", WxPay)
	wx.POST("/query", WxQuery)
	wx.POST("/reverse", WxReverse)
	wx.POST("/refund", WxRefund)
	wx.POST("/prepay", WxPrePay)
	wx.POST("/notify", WxNotify)

}

func InitDB(dialect, conn string) (newDb *xorm.Engine) {
	newDb, err := xorm.NewEngine(dialect, conn)
	if err != nil {
		panic(err)
	}
	return
}
