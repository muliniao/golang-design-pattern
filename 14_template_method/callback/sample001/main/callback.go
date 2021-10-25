package main

import "fmt"

// SuperCalculator ./
type SuperCalculator struct {

}

func (s *SuperCalculator) add(a, b int, xiaoming *Student)  {
	result := a + b
	xiaoming.FillBlank(a, b, result)
}


// Student ./
type Student struct {
	Name	string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (s *Student) CallHelp(a, b int) {
	new(SuperCalculator).add(a, b, s)
}

// FillBlank 回调函数 ./
func (s *Student) FillBlank(a, b, result int) {
	fmt.Printf(s.Name + "求助小红计算: [%v] + [%v] = [%v]", a, b, result)
}