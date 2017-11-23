package main

import (
	"fmt"
	"kit/test"
	"testing"
)

func Test_Account_Get(t *testing.T) {
	account, err := WxAccount{}.Get(10001)
	test.Ok(t, err)
	fmt.Println(account)
}
