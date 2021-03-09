package standard

import "fmt"

/**
具体Product001_002
*/
type ConcreteProduct001_002 struct {
}

func (concreteProduct001_002 *ConcreteProduct001_002) Show() {
	fmt.Println("具体产品001_002显示...")
}
