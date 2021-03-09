package main

import (
	"fmt"
	"golang-design-pattern/02_adapter/classadpter_standard"
)

func main() {

	fmt.Println("类适配器模式测试：")
	var target classadpter_standard.Target = new(classadpter_standard.ClassAdapter)
	target.Request()
}
