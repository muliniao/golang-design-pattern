package main

import (
	"fmt"
	"golang-design-pattern/11_command/standard"
)

func main() {

	var command standard.Command = standard.NewConcreteCommand(new(standard.Receiver))
	invoker := standard.NewInvoker(command)

	fmt.Println("Client访问调用者的call()方法!")
	invoker.Call()

}
