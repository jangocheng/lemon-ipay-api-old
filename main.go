package main

import (
	"flag"
	"lemon-ipay-api/datadb"
	"lemon-ipay-api/ipay"
	"lemon-ipay-api/wechat"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	appEnv      = flag.String("APP_ENV", os.Getenv("APP_ENV"), "APP_ENV")
	connEnv     = flag.String("IPAY_CONN", os.Getenv("IPAY_CONN"), "IPAY_CONN")
	bmappingUrl = flag.String("BMAPPING_URL", os.Getenv("BMAPPING_URL"), "BMAPPING_URL")
	envParam    ipay.EnvParamDto
)

func init() {
	flag.Parse()
	envParam = ipay.EnvParamDto{
		AppEnv:      *appEnv,
		ConnEnv:     *connEnv,
		BmappingUrl: *bmappingUrl,
	}
	datadb.Db = InitDB("mysql", envParam.ConnEnv)
	datadb.Db.Sync(new(datadb.Account), new(datadb.NotifyWechat))
}

func InitDB(dialect, conn string) (newDb *xorm.Engine) {
	newDb, err := xorm.NewEngine(dialect, conn)
	if err != nil {
		panic(err)
	}
	return
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
	v3.POST("/pay", ipay.Pay)
	v3.POST("/query", ipay.Query)
	v3.POST("/reverse", ipay.Reverse)
	v3.POST("/refund", ipay.Refund)
	v3.POST("/prepay", ipay.PrePay)

	v3.GET("/record/:Id", ipay.GetRecord)

	wx := v3.Group("/wx")
	wx.POST("/pay", wechat.Pay)
	wx.POST("/query", wechat.Query)
	wx.POST("/reverse", wechat.Reverse)
	wx.POST("/refund", wechat.Refund)
	wx.POST("/prepay", wechat.PrePay)
	wx.POST("/notify", wechat.Notify)

}
