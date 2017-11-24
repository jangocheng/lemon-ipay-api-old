package tests

import (
	"lemon-epay/datadb"
	"lemon-epay/ipay"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	envParam ipay.EnvParamDto
)

func init() {
	initTest()
	datadb.Db = InitDB("mysql", envParam.ConnEnv)
	datadb.Db.Sync(new(datadb.Account))
}

func initTest() {
	envParam = ipay.EnvParamDto{
		AppEnv:      "",
		ConnEnv:     os.Getenv("IPAY_CONN"),
		BmappingUrl: "",
	}
}

func InitDB(dialect, conn string) (newDb *xorm.Engine) {
	newDb, err := xorm.NewEngine(dialect, conn)
	if err != nil {
		panic(err)
	}
	return
}
