package standard002

type HandlerChain struct {
	List []IHandler
}

func NewHandlerChain() *HandlerChain {
	return &HandlerChain{
		List: make([]IHandler, 0),
	}
}

func (h *HandlerChain) AddHandler(handler IHandler) {
	h.List = append(h.List, handler)
}

func (h *HandlerChain) Handle() {
	for _, v := range h.List {
		if v.handler() {
			break
		}
	}
}