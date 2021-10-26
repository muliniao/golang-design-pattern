package standard

type Iterator interface {
	HasNext()		bool
	Next()			interface{}
	CurrentItem()	interface{}
}
