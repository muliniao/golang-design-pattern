package sample

type SortAlgFactory struct {
	Algs	map[string]ISortAlg
}

func NewSortAlgFactory() *SortAlgFactory {
	algs := make(map[string]ISortAlg)
	algs["QuickSort"] = new(QuickSort)
	algs["ExternalSort"] = new(ExternalSort)
	return &SortAlgFactory{
		Algs: algs,
	}
}

func (s *SortAlgFactory) GetSortAlg(sortType string) ISortAlg {
	return s.Algs[sortType]
}