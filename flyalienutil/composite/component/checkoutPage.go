package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"sync"
)

// CheckoutPage 订单结算页面组件
type CheckoutPage struct {
	composite.BaseComponent
}

func (bc *CheckoutPage) Do(ctx *composite.Context, currentConponent composite.Component, wg *sync.WaitGroup) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "订单结算页面组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}
