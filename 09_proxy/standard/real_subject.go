package standard

import "fmt"

/**
RealSubject: 真实主题
*/
type RealSubject struct {
}

func NewRealSubject() *RealSubject {
	return &RealSubject{}
}

func (realSubject *RealSubject) Request() {
	fmt.Println("访问真实主题方法...")
}
