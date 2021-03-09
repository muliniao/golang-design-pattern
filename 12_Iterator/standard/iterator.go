package standard

/**
Iterator: 抽象迭代器
*/
type Iterator interface {
	First() interface{}
	Next() interface{}
	HasNext() bool
}
