package static

import "fmt"

/**
Proxy: 代理
*/
type Proxy struct {
	realSubject *RealSubject
}

func NewProxy(realSubject *RealSubject) *Proxy {
	return &Proxy{
		realSubject: realSubject,
	}
}

func (proxy *Proxy) Request() {

	proxy.preRequest()
	if proxy.realSubject != nil {
		proxy.realSubject.Request()
	}
	proxy.postRequest()

}

func (proxy *Proxy) preRequest() {
	fmt.Println("访问真实主题之前的预处理...")

}

func (proxy *Proxy) postRequest() {
	fmt.Println("访问真实主题之后的后续处理...")
}
