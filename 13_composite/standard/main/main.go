package main

import "golang-design-pattern/13_composite/standard"

func main() {

	var c0 standard.Component = standard.NewComposite()
	var c1 standard.Component = standard.NewComposite()
	var leaf1 standard.Component = standard.NewLeaf("1")
	var leaf2 standard.Component = standard.NewLeaf("2")
	var leaf3 standard.Component = standard.NewLeaf("3")
	c0.Add(leaf1)
	c0.Add(c1)
	c1.Add(leaf2)
	c1.Add(leaf3)
	c0.Operation()
}
