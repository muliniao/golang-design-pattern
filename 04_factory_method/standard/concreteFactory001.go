package standard

import "fmt"

/*
	具体工厂001: 返回具体产品实例
 */
type ConcreteFactory001 struct {

}

func (ConcreteFactory001 *ConcreteFactory001) NewProduct() Product {
	fmt.Println("具体工厂001生成-->具体产品001...")
	return new(ConcreteProduct001)
}
