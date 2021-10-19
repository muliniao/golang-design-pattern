package main

import (
	"fmt"

	"golang-design-pattern/20_decorator/standard"
)

func main() {

	/*
		(1) 第一个比较特殊的地方是：装饰器类和原始类继承同样的父类，这样我们可以对原始类“嵌套”多个装饰器类。
		(2) 第二个比较特殊的地方是：装饰器类是对功能的增强，这也是装饰器模式应用场景的一个重要特点。
			- 代理模式: 附加新功能,与原始类无关
			- 装饰器模式: 增强功能,与原始类有关

	 */

	var component001 standard.Component = standard.NewConcreteComponent()
	component001.Operation()

	fmt.Println("-----------------------------------------------------")

	var component002 standard.Component = standard.NewConcreteDecorator(standard.NewBaseDecorator(component001))
	component002.Operation()
}
