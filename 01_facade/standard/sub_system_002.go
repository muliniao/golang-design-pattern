package standard

import "fmt"

/**
SubSystem002: 子系统
*/
type SubSystem002 struct {
}

func NewSubSystem002() *SubSystem002 {
	return &SubSystem002{}
}

func (subSystem002 *SubSystem002) method002() {
	fmt.Println("子系统002的method()被调用")
}
