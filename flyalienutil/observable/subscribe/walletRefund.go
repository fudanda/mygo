package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// WalletRefund 退钱包余额
type WalletRefund struct{}

// Do 具体业务
func (w *WalletRefund) Do(o observable.Observable) (err error) {
	// code...
	fmt.Println(util.RunFuncName(), "退钱包余额...")
	return
}
