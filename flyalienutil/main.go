package main

import (
	"flyalienutil/composite"
	"flyalienutil/composite/component"
	"flyalienutil/responsibility"
	"flyalienutil/responsibility/handler"
	"fmt"
)

func main() {

	fmt.Println("responsibility | 责任链 | --------")
	// 初始化空handler
	nullHandler := &responsibility.NullHandler{}

	// 链式调用 代码是不是很优雅
	// 很明显的链 逻辑关系一览无余
	nullHandler.SetNext(&responsibility.ArgumentsHandler{}).
		SetNext(&handler.AddressInfo{}).
		SetNext(&handler.StockInfo{})
		//无限扩展代码...

	// 开始执行业务
	if err := nullHandler.Run(&responsibility.Context{}); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}

	//组合模式
	fmt.Println("component | 组合模式 | --------")

	checkoutPage := &component.CheckoutPage{}

	skuComponent := &component.Sku{}
	orderComponent := &component.Order{}

	skuComponent.Mount(
		&component.Promotion{},
		&component.AfterSale{},
	)

	checkoutPage.Mount(
		&component.Address{},
		skuComponent,
		orderComponent,
	)

	checkoutPage.Do(&composite.Context{})

	checkoutPage.Remove(orderComponent)
	checkoutPage.Do(&composite.Context{})

	// 成功
	fmt.Println("Success")
}
