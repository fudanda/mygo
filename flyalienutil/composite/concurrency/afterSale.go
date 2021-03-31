package concurrency

import (
	"flyalienutil/composite"
	"flyalienutil/util"
	"fmt"
	"net/http"
)

// AfterSale 售后组件
type AfterSale struct {
	composite.BaseConcurrencyComponent
}

func (bc *AfterSale) BusinessLogicDo(resChan chan interface{}) error {
	fmt.Println(util.RunFuncName(), "售后组件...")
	fmt.Println(util.RunFuncName(), "获取售后信息 ing...")

	// 模拟远程调用地址服务
	http.Get("https://www.baidu.com/")

	resChan <- struct{}{} // 写入业务执行结果
	fmt.Println(util.RunFuncName(), "获取售后信息 done...")
	return nil
}
