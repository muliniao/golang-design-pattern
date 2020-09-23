package main

import (
	"fmt"
	"golang-design-pattern/17_memento/standard"
)

func main() {

	originator := new(standard.Originator)
	caretaker := new(standard.Caretaker)

	originator.SetState("No1")
	fmt.Printf("初始状态: [%s]\n", originator.GetState())

	// 保存状态
	caretaker.SetMemento(originator.CreateMemento())

	//更新新的状态
	originator.SetState("No2")
	fmt.Printf("新的状态: [%s]\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento())

	fmt.Printf("恢复状态: [%s]\n", originator.GetState())

}
