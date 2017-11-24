package datadb

import (
	"fmt"
	"kit/test"
	"testing"
)

func Test_Account_Get(t *testing.T) {
	account, err := Account{}.Get(10001)
	test.Ok(t, err)
	fmt.Println(account)
}
