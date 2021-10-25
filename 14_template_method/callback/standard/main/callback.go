package main

import "fmt"

type ICallback interface {
	methodTOCallback()
}

// ClassB ./
type ClassB struct {

}

func (b *ClassB) Process(callback ICallback) {
	fmt.Println("ClassB: process")
	callback.methodTOCallback()
}

// ClassA ./
type ClassA struct {

}

func (a *ClassA) methodTOCallback() {
	fmt.Println("ClassA: callback")
}

func (a *ClassA) Process() {
	new(ClassB).Process(new(ClassA))
}


func main() {
	a :=  new(ClassA)
	a.Process()
}