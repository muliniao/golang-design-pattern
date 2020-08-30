package standard

import "fmt"

/**
	具体Product001_001
*/
type ConcreteProduct001_001 struct {

}

func (concreteProduct001_001 *ConcreteProduct001_001) Show() {
	fmt.Println("具体产品001_001显示...")
}
