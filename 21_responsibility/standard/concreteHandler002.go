package standard

import "fmt"

/**
ConcreteHandler002: 具体处理者
*/
type ConcreteHandler002 struct {
	AbstractHandler AbstractHandler
}

func (c *ConcreteHandler002) HandleRequest(request string) {

	if request == "two" {
		fmt.Println("ConcreteHandler002 deal with this request")
	} else {
		if c.AbstractHandler.GetNext() != nil {
			c.AbstractHandler.GetNext().HandleRequest(request)
		} else {
			fmt.Println("No one handle")
		}
	}
}
