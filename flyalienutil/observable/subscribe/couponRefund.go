package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// CouponRefund 退优惠券
type CouponRefund struct{}

func (c *CouponRefund) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "退优惠券...")
	return
}
