package concurrency

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
)

// Order 订单详细信息组件
type Order struct {
	composite.BaseComponent
}

func (bc *Order) BusinessLogicDo(resChan chan interface{}) (err error) {
	fmt.Println(util.RunFuncName(), "订单详细信息组件...")

	return
}
