package main

import "golang-design-pattern/14_template_method/standard"

func main() {

	concreteTemplate := new(standard.ConcreteTemplate)
	template := standard.NewTemplateFather(concreteTemplate)
	template.Result()

}
