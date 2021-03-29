package subscribe

import (
	"flyalienutil/observable"
	"flyalienutil/util"
	"fmt"
)

// Sms 发短信
type Sms struct{}

func (s *Sms) Do(o observable.Observable) (err error) {
	fmt.Println(util.RunFuncName(), "发短信...")
	return
}
