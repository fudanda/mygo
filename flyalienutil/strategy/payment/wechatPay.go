package payment

import (
	"flyalienutil/strategy"
	"flyalienutil/util"
	"fmt"
)

// WechatPay 使用微信支付
type WechatPay struct{}

func (w *WechatPay) Pay(ctx *strategy.Context) (err error) {
	fmt.Println(util.RunFuncName(), "使用微信支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (w *WechatPay) Refund(ctx *strategy.Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(util.RunFuncName(), "使用微信退款...")
	return
}
