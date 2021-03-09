package standard

import "fmt"

/*
	具体产品002
*/
type ConcreteProduct002 struct {
}

func (concreteProduct002 *ConcreteProduct002) Show() {
	fmt.Println("具体产品002显示...")
}
