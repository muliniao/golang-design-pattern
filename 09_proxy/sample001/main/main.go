package main

import "golang-design-pattern/09_proxy/sample001"

func main()  {

	userProxy := sample001.NewUserControllerProxy(new(sample001.UserController))
	userProxy.Register("12345678901", "abcd1234")

}
