package classadpter_standard

import "fmt"

/**
被适配者
*/
type Adaptee struct {
}

func (adaptee *Adaptee) SpecificRequest() {
	fmt.Println("[类适配]被适配者中的业务代码被调用！")
}
