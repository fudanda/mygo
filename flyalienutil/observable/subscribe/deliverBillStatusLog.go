package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// DeliverBillStatusLog 记录发货单状态变更日志
type DeliverBillStatusLog struct{}

func (m *DeliverBillStatusLog) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "记录发货单状态变更日志...")
	return
}
