package handler

import (
	"flyalienutil/responsibility"
	"flyalienutil/util"
	"fmt"
)

type AddressInfo struct {
	responsibility.Next
}

// Do 校验参数的逻辑
func (h *AddressInfo) Do(c *responsibility.Context) (err error) {
	fmt.Println(util.RunFuncName(), "获取地址信息...")
	fmt.Println(util.RunFuncName(), "地址信息校验...")
	return
}
