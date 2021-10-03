package standard002

import "fmt"

type ConcreteHandler002 struct {

}

func (c *ConcreteHandler002) handler() bool {
	handler := false
	fmt.Println("concreteHandler002: false")
	return handler
}
