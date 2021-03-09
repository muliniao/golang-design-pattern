package standard

import "fmt"

/*
	具体工厂002: 返回具体产品实例
*/
type ConcreteFactory002 struct {
}

func (concreteFactory002 *ConcreteFactory002) NewProduct() Product {
	fmt.Println("具体工厂002生成-->具体产品002...")
	return new(ConcreteProduct002)
}
