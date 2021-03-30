package payment

import (
	"flyalienutil/strategy"
	"flyalienutil/util"
	"fmt"
)

// BankPay 银行卡支付
type BankPay struct{}

// Pay 当前支付方式的支付逻辑
func (p *BankPay) Pay(ctx *strategy.Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(util.RunFuncName(), "使用银行卡支付...")
	return
}

// Refund 当前支付方式的支付逻辑
func (p *BankPay) Refund(ctx *strategy.Context) (err error) {
	// 当前策略的业务逻辑写这
	fmt.Println(util.RunFuncName(), "使用银行卡退款...")
	return
}
