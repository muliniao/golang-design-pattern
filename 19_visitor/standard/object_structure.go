package standard

/**
	Object Structure
 */
type ObjectStructure struct {
	objectList	[]Element
}

func NewObjectStructure() *ObjectStructure {
	return &ObjectStructure{
		objectList: make([]Element, 0),
	}
}

func (objectStructure *ObjectStructure) Attach(element Element) {

	objectStructure.objectList = append(objectStructure.objectList, element)

}

func (objectStructure *ObjectStructure) Display(visitor Visitor) {

	for _,v := range objectStructure.objectList {
		v.Accept(visitor)
	}
}
