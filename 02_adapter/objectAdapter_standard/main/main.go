package main

import (
	"fmt"
	"golang-design-pattern/02_adapter/objectAdapter_standard"
)

func main() {

	fmt.Println("对象适配器模式测试：")
	adaptee := objectAdapter_standard.NewAdaptee()
	var target objectAdapter_standard.Target = objectAdapter_standard.NewObjectAdapter(adaptee)
	target.Request()

}
