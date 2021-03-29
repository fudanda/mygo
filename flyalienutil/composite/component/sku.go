package component

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
)

// Sku 商品信息组件
type Sku struct {
	composite.BaseComponent
}

func (bc *Sku) Do(ctx *composite.Context) (err error) {
	// 当前组件的业务逻辑写这
	fmt.Println(util.RunFuncName(), "商品信息组件...")

	bc.ChildsDo(ctx)

	return
}
