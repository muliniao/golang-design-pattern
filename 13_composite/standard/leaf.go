package standard

import "fmt"

/**
	Leaf:叶子节点，最底层
 */
type Leaf struct {
	Name	string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{
		Name: name,
	}
}

func (leaf *Leaf) Add(component Component) {

}

func (leaf *Leaf) Remove(component Component) {

}

func (leaf *Leaf) GetChild(i int) Component {
	return nil
}

func (leaf *Leaf) Operation() {
	fmt.Printf("树叶[%s]:被访问", leaf.Name)
}

