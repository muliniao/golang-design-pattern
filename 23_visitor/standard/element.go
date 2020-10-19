package standard

/**
	Element
 */
type Element interface {
	Accept(visitor Visitor)
}
