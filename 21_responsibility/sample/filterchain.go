package sample

import "fmt"

type FilterChain struct {
	position	int
	total		int
	filters     []SensitiveWordFilter
}

func NewFilterChain() *FilterChain {
	return &FilterChain{
		position: 0,
		total:    0,
		filters:  make([]SensitiveWordFilter, 0),
	}
}

func (f *FilterChain) AddFilter(sensitiveWordFilter SensitiveWordFilter) {
	f.filters = append(f.filters, sensitiveWordFilter)
	f.total ++
}

func (f *FilterChain) DoFilter(content string) {
	if f.position < f.total {
		filter := f.filters[f.position]
		f.position++
		if filter.doFilter(content) {
			return
		}

		f.DoFilter(content)
	} else {
		fmt.Println("通过敏感词filter")
	}
}
