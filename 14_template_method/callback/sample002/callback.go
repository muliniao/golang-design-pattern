package sample002

import "fmt"

type DoJob interface {
	fillBlank(a, b, result int)
}

// SuperCalculator ./
type SuperCalculator struct {

}

func (s *SuperCalculator) add(a, b int, customer DoJob) {
	result := a + b
	customer.fillBlank(a, b, result)
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

// fillBlank: callBack() ./
func (s *Student) fillBlank(a, b, result int) {
	fmt.Printf(s.Name + "求助小红计算: [%v] + [%v] = [%v]", a, b, result)
}

func (s *Student) CanHelp(a, b int) {
	new(SuperCalculator).add(a, b, s)
}