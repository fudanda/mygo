package main

import (
	"fmt"
	"runtime"
)

//------------------------------------------------------------
//设计模式
//模板模式
//------------------------------------------------------------
// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

const (
	// ConstActTypeTime 按时间类型抽奖
	ConstActTypeTime int = 1
	// ConstActTypeTimes 按次数抽奖
	ConstActTypeTimes int = 2
	// ConstActTypeAmount 按范围区间抽奖
	ConstActTypeAmount int = 3
)

// Context上下文
type Context struct {
	ActInfo *ActInfo
}
type ActInfo struct {
	//抽奖类型
	ActivityType int
}

// BehaviorInterface 不同抽奖类型的行为差异的抽象接口
type BehaviorInterface interface {
	// checkParams 其他参数校验
	checkParams(ctx *Context) error
	// getPrizesByNode 获取奖品信息
	getPrizesByNode(ctx *Context) error
}

// 时间类抽奖类型
type TimeDraw struct {
	// 合成复用模板
	Lottery
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw TimeDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "do nothing(抽取该场次的奖品即可，无需其他逻辑)...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
func (draw TimeDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "do nothing(抽取该场次的奖品即可，无需其他逻辑)...")
	return
}

// TimesDraw 具体抽奖行为
// 按抽奖次数抽奖类型 比如答题闯关
type TimesDraw struct {
	// 合成复用模板
	Lottery
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw TimesDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "按抽奖次数抽奖类型:特殊参数校验...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
func (draw TimesDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "1. 判断是该用户第几次抽奖...")
	fmt.Println(runFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(runFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return
}

// AmountDraw 具体抽奖行为
// 按数额范围区间抽奖 比如订单金额刮奖
type AmountDraw struct {
	// 合成复用模板
	Lottery
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw *AmountDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "按数额范围区间抽奖:特殊参数校验...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
func (draw *AmountDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "1. 判断属于哪个数额区间...")
	fmt.Println(runFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(runFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return
}

// Lottery 抽奖模板
type Lottery struct {
	// 不同抽奖类型的抽象行为
	concreteBehavior BehaviorInterface
}

// Run 抽奖算法
// 稳定不变的算法步骤
func (lottery *Lottery) Run(ctx *Context) (err error) {
	// 具体方法：校验活动编号(serial_no)是否存在、并获取活动信息
	if err = lottery.checkSerialNo(ctx); err != nil {
		return err
	}

	// 具体方法：校验活动、场次是否正在进行
	if err = lottery.checkStatus(ctx); err != nil {
		return err
	}

	// ”抽象方法“：其他参数校验
	if err = lottery.checkParams(ctx); err != nil {
		return err
	}

	// 具体方法：活动抽奖次数校验(同时扣减)
	if err = lottery.checkTimesByAct(ctx); err != nil {
		return err
	}

	// 具体方法：活动是否需要消费积分
	if err = lottery.consumePointsByAct(ctx); err != nil {
		return err
	}

	// 具体方法：场次抽奖次数校验(同时扣减)
	if err = lottery.checkTimesBySession(ctx); err != nil {
		return err
	}

	// 具体方法：获取场次奖品信息
	if err = lottery.getPrizesBySession(ctx); err != nil {
		return err
	}

	// ”抽象方法“：获取node奖品信息
	if err = lottery.getPrizesByNode(ctx); err != nil {
		return err
	}

	// 具体方法：抽奖
	if err = lottery.drawPrizes(ctx); err != nil {
		return err
	}

	// 具体方法：奖品数量判断
	if err = lottery.checkPrizesStock(ctx); err != nil {
		return err
	}

	// 具体方法：组装奖品信息
	if err = lottery.packagePrizeInfo(ctx); err != nil {
		return err
	}
	return
}

// checkSerialNo 校验活动编号(serial_no)是否存在
func (lottery *Lottery) checkSerialNo(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "校验活动编号(serial_no)是否存在、并获取活动信息...")
	return
}

// checkStatus 校验活动、场次是否正在进行
func (lottery *Lottery) checkStatus(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "校验活动、场次是否正在进行...")
	return
}

// checkParams 其他参数校验(不同活动类型实现不同)
// 不同场景变化的算法 转化为依赖抽象
func (lottery *Lottery) checkParams(ctx *Context) (err error) {
	// 实际依赖的接口的抽象方法
	return lottery.concreteBehavior.checkParams(ctx)
}

// checkTimesByAct 活动抽奖次数校验
func (lottery *Lottery) checkTimesByAct(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "活动抽奖次数校验...")
	return
}

// consumePointsByAct 活动是否需要消费积分
func (lottery *Lottery) consumePointsByAct(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "活动是否需要消费积分...")
	return
}

// checkTimesBySession 活动抽奖次数校验
func (lottery *Lottery) checkTimesBySession(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "活动抽奖次数校验...")
	return
}

// getPrizesBySession 获取场次奖品信息
func (lottery *Lottery) getPrizesBySession(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "获取场次奖品信息...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
// 不同场景变化的算法 转化为依赖抽象
func (lottery *Lottery) getPrizesByNode(ctx *Context) (err error) {
	// 实际依赖的接口的抽象方法
	return lottery.concreteBehavior.getPrizesByNode(ctx)
}

// drawPrizes 抽奖
func (lottery *Lottery) drawPrizes(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "抽奖...")
	return
}

// checkPrizesStock 奖品数量判断
func (lottery *Lottery) checkPrizesStock(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "奖品数量判断...")
	return
}

// packagePrizeInfo 组装奖品信息
func (lottery *Lottery) packagePrizeInfo(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "组装奖品信息...")
	return
}

func main() {
	ctx := &Context{
		ActInfo: &ActInfo{
			ActivityType: ConstActTypeAmount,
		},
	}

	switch ctx.ActInfo.ActivityType {
	case ConstActTypeTime: // 按时间抽奖类型
		instance := &TimeDraw{}
		instance.concreteBehavior = instance
		instance.Run(ctx)
	case ConstActTypeTimes: // 按抽奖次数抽奖
		instance := &TimesDraw{}
		instance.concreteBehavior = instance
		instance.Run(ctx)
	case ConstActTypeAmount: // 按数额范围区间抽奖
		instance := &AmountDraw{}
		instance.concreteBehavior = instance
		instance.Run(ctx)
	default:
		// 报错
		return
	}
}
