package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"sync"
)

// AfterSale 售后组件
type AfterSale struct {
	composite.BaseComponent
}

func (bc *AfterSale) Do(ctx *composite.Context, currentConponent composite.Component, wg *sync.WaitGroup) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "售后组件...")

	// 执行子组件
	bc.ChildsDo(ctx)

	// 当前组件的业务逻辑写这

	return
}
