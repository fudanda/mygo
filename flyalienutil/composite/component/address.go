package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
)

// Address 地址信息组件
type Address struct {
	composite.BaseComponent
}

func (bc *Address) Do(ctx *composite.Context) (err error) {
	fmt.Println(util.RunFuncName(), "地址信息组件...")

	bc.ChildsDo(ctx)

	return
}
