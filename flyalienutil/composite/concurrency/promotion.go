package concurrency

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
)

// Promotion 优惠信息组件
type Promotion struct {
	composite.BaseComponent
}

func (bc *Promotion) BusinessLogicDo(resChan chan interface{}) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "优惠信息组件...")

	return

}
