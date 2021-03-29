package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// Refund 生成退款单
type Refund struct{}

func (refund *Refund) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "生成退款单...")
	return
}
