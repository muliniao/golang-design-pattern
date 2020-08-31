package standard

import "fmt"

/**
	SubSystem001: 子系统
 */
type SubSystem001 struct {
	
}

func NewSubSystem001() *SubSystem001 {
	return &SubSystem001{}
}

func (subSystem001 *SubSystem001) method001() {
	fmt.Println("子系统001的method()被调用")
}
