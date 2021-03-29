package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// PromotionRefund 还优惠活动资格
type PromotionRefund struct{}

func (p *PromotionRefund) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "还优惠活动资格...")
	return
}
