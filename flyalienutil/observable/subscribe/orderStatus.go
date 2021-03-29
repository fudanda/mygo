package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// OrderStatus 修改订单状态
type OrderStatus struct{}

func (observer *OrderStatus) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "修改订单状态...")
	return
}
