package standard

import "fmt"

/**
	Receiver: 实现者/接收者角色
 */
type Receiver struct {

}

func (receiver *Receiver) Action() {
	fmt.Println("接收者的action()方法被调用!")
}
