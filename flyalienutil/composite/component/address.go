package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"sync"
)

// Address 地址信息组件
type Address struct {
	composite.BaseComponent
}

func (bc *Address) Do(ctx *composite.Context, currentConponent composite.Component, wg *sync.WaitGroup) (err error) {
	fmt.Println(util.RunFuncName(), "地址信息组件...")

	bc.ChildsDo(ctx)

	return
}
