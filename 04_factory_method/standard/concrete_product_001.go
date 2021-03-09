package standard

import "fmt"

/*
	具体产品001
*/
type ConcreteProduct001 struct {
}

func (concreteProduct001 *ConcreteProduct001) Show() {
	fmt.Println("具体产品001显示...")
}
