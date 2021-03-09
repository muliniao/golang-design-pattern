package standard

/**
Subject: 抽象主题
*/
type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObserver()
}
