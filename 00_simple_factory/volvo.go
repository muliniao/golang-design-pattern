package _0_simple_factory

import "fmt"

/**
产品
*/
type Volvo struct {
}

func (volvo *Volvo) Operation() {
	fmt.Println("This is Volvo")
}
