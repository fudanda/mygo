package concurrency

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"net/http"
)

// Address 地址信息组件
type Address struct {
	composite.BaseConcurrencyComponent
}

func (bc *Address) BusinessLogicDo(resChan chan interface{}) error {
	fmt.Println(util.RunFuncName(), "地址组件...")
	fmt.Println(util.RunFuncName(), "获取地址信息 ing...")

	// 模拟远程调用地址服务
	http.Get("https://www.baidu.com/")

	resChan <- struct{}{} // 写入业务执行结果
	fmt.Println(util.RunFuncName(), "获取地址信息 done...")
	return nil
}
