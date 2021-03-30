package payment

import (
	"flyalienutil/strategy"
	"flyalienutil/util"
	"fmt"
)

// AliPayWap 使用支付宝网页版
type AliPayWap struct{}

func (a *AliPayWap) Pay(ctx *strategy.Context) (err error) {
	fmt.Println(util.RunFuncName(), "使用支付宝网页版支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (a *AliPayWap) Refund(ctx *strategy.Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(util.RunFuncName(), "使用支付宝网页版退款...")
	return
}
