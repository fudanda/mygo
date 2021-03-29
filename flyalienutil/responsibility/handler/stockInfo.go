package handler

import (
	"flyalienutil/responsibility"
	"flyalienutil/util"
	"fmt"
)

const (
	EMPTY     string = "商品-不存在"
	UNALLOWED string = "校验失败-库存不足"
)

type StockInfo struct {
	responsibility.Next
}

// Do 校验参数的逻辑
func (s *StockInfo) Do(c *responsibility.Context) (err error) {
	fmt.Println(util.RunFuncName(), "获取库存信息...")
	fmt.Println(util.RunFuncName(), "库存信息校验...")
	//手动抛出异常
	// f := -1
	// if f < 0 {
	// 	return errors.New(UNALLOWED)
	// }
	return
}
