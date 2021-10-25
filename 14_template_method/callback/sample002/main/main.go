package main

import "golang-design-pattern/14_template_method/callback/sample002"

func main() {

	s := sample002.NewStudent("小明")
	s.CanHelp(12,13)

}
