package composite

import (
	"flyalienutil/util"
	"fmt"
	"reflect"
)

//------------------------------------------------------------
//组合模式
//------------------------------------------------------------

// Context 上下文
type Context struct{}

type Component interface {
	//Mount 添加组件
	Mount(c Component, components ...Component) error
	//Remove 删除组件
	Remove(c Component) error
	//Do 执行
	Do(ctx *Context) error
}

// BaseComponent 基础组件
// 实现Mount:添加一个子组件
// 实现Remove:移除一个子组件
type BaseComponent struct {
	ChildComponents []Component
}

func (bc *BaseComponent) Mount(c Component, components ...Component) (err error) {
	bc.ChildComponents = append(bc.ChildComponents, c)
	if len(components) == 0 {
		return
	}
	bc.ChildComponents = append(bc.ChildComponents, components...)
	return
}

func (bc *BaseComponent) Remove(c Component) (err error) {
	if len(bc.ChildComponents) == 0 {
		return
	}
	for k, childComponent := range bc.ChildComponents {
		if c == childComponent {
			fmt.Println(util.RunFuncName(), "移除:", reflect.TypeOf(childComponent))
			bc.ChildComponents = append(bc.ChildComponents[:k], bc.ChildComponents[k+1:]...)
		}
	}
	return
}

func (bc *BaseComponent) Do(ctx *Context) (err error) {
	return
}

func (bc *BaseComponent) ChildsDo(ctx *Context) (err error) {
	for _, childComponent := range bc.ChildComponents {
		if err = childComponent.Do(ctx); err != nil {
			return err
		}
	}
	return
}
