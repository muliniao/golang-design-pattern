package sample

type ISortAlg interface {
	sort(filePath string)
}

// QuickSort ./
type QuickSort struct {

}

func (q *QuickSort) sort(filePath string) {

}

// ExternalSort ./
type ExternalSort struct {

}

func (e *ExternalSort) sort(filePath string) {

}