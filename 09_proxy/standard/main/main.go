package main

import "golang-design-pattern/09_proxy/standard"

func main() {

	proxy := standard.NewProxy(standard.NewRealSubject())
	proxy.Request()

}
