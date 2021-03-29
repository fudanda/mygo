package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// DeliverBillStatus 修改发货单状态
type DeliverBillStatus struct{}

func (m *DeliverBillStatus) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "修改发货单状态...")
	return
}
