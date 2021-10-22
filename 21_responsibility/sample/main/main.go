package main

import "golang-design-pattern/21_responsibility/sample"

func main() {

	filterChain := sample.NewFilterChain()
	filterChain.AddFilter(new(sample.SexyWordFilter))
	filterChain.AddFilter(new(sample.PoliticalWordFilter))

	filterChain.DoFilter("习大大")

}
