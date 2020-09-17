package standard

/**
	ConcreteSubject: 具体主题
 */
type ConcreteSubject struct {
	ObserverList	[]Observer
}

func NewConcreteSubject() *ConcreteSubject {
 	return &ConcreteSubject{
		ObserverList: make([]Observer, 0),
	}
}

func (concreteSubject *ConcreteSubject) AddObserver(observer Observer) {
	concreteSubject.ObserverList = append(concreteSubject.ObserverList, observer)
}

func (concreteSubject *ConcreteSubject) RemoveObserver(observer Observer) {
	for key, value := range concreteSubject.ObserverList {
		if value == observer {
			newObserverList := append(concreteSubject.ObserverList[:key], concreteSubject.ObserverList[key+1:]...)
			concreteSubject.ObserverList = newObserverList
		}
	}
}

func (concreteSubject *ConcreteSubject) NotifyObserver() {
	for _, value := range concreteSubject.ObserverList {
		a := value
		a.Response()
	}
}
