package sample

import (
	"fmt"
	"strings"
)

type SexyWordFilter struct {
}

func (s *SexyWordFilter) doFilter(content string) bool {
	fmt.Println("sexy filter")
	return strings.Contains(content, "性")
}
