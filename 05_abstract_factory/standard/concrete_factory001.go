package standard

import "fmt"

/**
	具体工厂001
 */
type ConcreteFactory001 struct {

}

func (concreteFactory001 *ConcreteFactory001) NewProduct001() Product001 {
	fmt.Println("具体工厂001生产 ---> 具体产品001_001")
	return new(ConcreteProduct001_001)
}

func (concreteFactory001 *ConcreteFactory001) NewProduct002() Product002 {
	fmt.Println("具体工厂001生产 ---> 具体产品002_001")
	return new(ConcreteProduct002_001)
}
