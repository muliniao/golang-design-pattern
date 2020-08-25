package _0_simple_factory

import "fmt"

/**
	产品
 */
type Audio struct {

}

func (audio *Audio) Operation() {
	fmt.Println("This is Audi")
}

