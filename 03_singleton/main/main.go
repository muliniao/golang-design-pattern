package main

import (
	"fmt"
	singleton "golang-design-pattern/03_singleton"
)

func main() {

	instance001 := singleton.GetInstance()
	instance002 := singleton.GetInstance()

	if instance001 != instance002 {
		fmt.Println("instance is not equal")
	} else {
		fmt.Println("instance is equal")
	}
}
