package standard001

/**
Handler: 抽象处理者
*/
type Handler interface {
	HandleRequest(request string)
}

type AbstractHandler struct {
	Next Handler
}

func (a *AbstractHandler) SetNext(next Handler) {
	a.Next = next
}

func (a *AbstractHandler) GetNext() Handler {
	return a.Next
}
