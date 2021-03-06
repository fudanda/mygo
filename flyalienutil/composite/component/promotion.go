package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"sync"
)

// Promotion 优惠信息组件
type Promotion struct {
	composite.BaseComponent
}

func (bc *Promotion) Do(ctx *composite.Context, currentConponent composite.Component, wg *sync.WaitGroup) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "优惠信息组件...")

	bc.ChildsDo(ctx)

	return

}
