package main

import (
	"golang-design-pattern/09_proxy/standard/static"
)

func main() {

	proxy := static.NewProxy(static.NewRealSubject())
	proxy.Request()

}
