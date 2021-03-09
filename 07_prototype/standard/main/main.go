package main

import (
	"fmt"
	"golang-design-pattern/07_prototype/standard"
)

func main() {

	protoManger := standard.NewPrototypeManager()

	realizeType := &standard.RealizeType{
		Name: "Rocket",
		Age:  34,
	}

	protoManger.Set("realizeType001", realizeType)
	c := protoManger.Get("realizeType001")

	realizeType001 := c.(*standard.RealizeType)

	fmt.Printf("[realizeType001.name]: [%s]", realizeType001.Name)
	fmt.Printf("[realizeType001.age]: [%v]", realizeType001.Age)

}
