package standard002

import "fmt"

type ConcreteHandler001 struct {

}

func (c *ConcreteHandler001) handler() bool {
	handler := false
	fmt.Println("concreteHandler001: false")
	return handler
}
