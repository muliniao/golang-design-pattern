package objectAdapter_standard

import "fmt"

/**
被适配者
*/
type Adaptee struct {
}

func NewAdaptee() *Adaptee {
	return &Adaptee{}
}

func (adaptee *Adaptee) SpecificRequest() {
	fmt.Println("[对象适配]被适配者中的业务代码被调用！")
}
