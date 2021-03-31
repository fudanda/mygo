package concurrency

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
)

// CheckoutPage 订单结算页面组件
type CheckoutPage struct {
	composite.BaseConcurrencyComponent
}

func (bc *CheckoutPage) BusinessLogicDo(resChan chan interface{}) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "订单结算页面组件...")

	// 当前组件的业务逻辑写这

	return
}
