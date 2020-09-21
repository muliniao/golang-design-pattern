package standard

/**
	Aggregate: 抽象聚合
 */
type Aggregate interface {
	Add(Object interface{})
	Remove(Object interface{})
	GetIterator() Iterator
}
