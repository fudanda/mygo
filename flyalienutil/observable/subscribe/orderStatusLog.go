package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// OrderStatusLog 记录订单状态变更日志
type OrderStatusLog struct{}

func (l *OrderStatusLog) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "记录订单状态变更日志...")
	return
}
