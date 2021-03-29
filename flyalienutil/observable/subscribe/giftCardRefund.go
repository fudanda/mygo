package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// GiftCardRefund 返礼品卡
type GiftCardRefund struct{}

func (r *GiftCardRefund) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "返礼品卡...")
	return
}
