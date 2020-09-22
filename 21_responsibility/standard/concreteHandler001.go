package standard

import "fmt"

/**
	ConcreteHandler001: 具体处理者
 */
type ConcreteHandler001 struct {
	AbstractHandler	AbstractHandler
}

func (c *ConcreteHandler001) HandleRequest(request string) {

	if request == "one" {
		fmt.Println("ConcreteHandler001 deal with this request")
	} else {
		if c.AbstractHandler.GetNext() != nil {
			c.AbstractHandler.GetNext().HandleRequest(request)
		} else {
			fmt.Println("No one handle")
		}
	}
}
