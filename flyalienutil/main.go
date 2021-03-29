package main

import (
	"flyalienutil/composite"
	"flyalienutil/composite/component"
	"flyalienutil/observable"
	"flyalienutil/observable/subscribe"
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

	//观察者模式
	fmt.Println("observable | 观察者模式 | --------")

	// 创建 未支付取消订单 “主题”
	fmt.Println("----------------------- 未支付取消订单 “主题”")
	orderUnPaidCancelSubject := &observable.ObservableConcrete{}
	orderUnPaidCancelSubject.Attach(
		&subscribe.OrderStatus{},
		&subscribe.OrderStatusLog{},
		&subscribe.CouponRefund{},
		&subscribe.PromotionRefund{},
		&subscribe.StockRefund{},
	)
	orderUnPaidCancelSubject.Notify()

	// 创建 超时关单 “主题”
	fmt.Println("----------------------- 超时关单 “主题”")
	orderOverTimeSubject := &observable.ObservableConcrete{}
	orderOverTimeSubject.Attach(
		&subscribe.OrderStatus{},
		&subscribe.OrderStatusLog{},
		&subscribe.CouponRefund{},
		&subscribe.PromotionRefund{},
		&subscribe.StockRefund{},
		&subscribe.Sms{},
	)
	orderOverTimeSubject.Notify()

	// 创建 已支付取消订单 “主题”
	fmt.Println("----------------------- 已支付取消订单 “主题”")
	orderPaidCancelSubject := &observable.ObservableConcrete{}
	orderPaidCancelSubject.Attach(
		&subscribe.OrderStatus{},
		&subscribe.OrderStatusLog{},
		&subscribe.CouponRefund{},
		&subscribe.PromotionRefund{},
		&subscribe.StockRefund{},
		&subscribe.GiftCardRefund{},
		&subscribe.WalletRefund{},
		&subscribe.Refund{},
		&subscribe.Sms{},
	)
	orderPaidCancelSubject.Notify()

	// 创建 取消发货单 “主题”
	fmt.Println("----------------------- 取消发货单 “主题”")
	deliverBillCancelSubject := &observable.ObservableConcrete{}
	deliverBillCancelSubject.Attach(
		&subscribe.OrderStatus{},
		&subscribe.OrderStatusLog{},
		&subscribe.DeliverBillStatus{},
		&subscribe.DeliverBillStatusLog{},
		&subscribe.StockRefund{},
		&subscribe.GiftCardRefund{},
		&subscribe.WalletRefund{},
		&subscribe.Refund{},
		&subscribe.Sms{},
	)
	deliverBillCancelSubject.Notify()

	// 创建 拒收 “主题”
	fmt.Println("----------------------- 拒收 “主题”")
	deliverBillRejectSubject := &observable.ObservableConcrete{}
	deliverBillRejectSubject.Attach(
		&subscribe.OrderStatus{},
		&subscribe.OrderStatusLog{},
		&subscribe.DeliverBillStatus{},
		&subscribe.DeliverBillStatusLog{},
		&subscribe.StockRefund{},
		&subscribe.GiftCardRefund{},
		&subscribe.WalletRefund{},
		&subscribe.Refund{},
		&subscribe.Sms{},
	)
	deliverBillRejectSubject.Notify()

	// 成功
	fmt.Println("Success")
}
