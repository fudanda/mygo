package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// StockRefund 还库存
type StockRefund struct{}

func (s *StockRefund) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "还库存...")
	return
}
