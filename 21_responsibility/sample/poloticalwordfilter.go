package sample

import (
	"fmt"
	"strings"
)

type PoliticalWordFilter struct {

}

func (p *PoliticalWordFilter) doFilter(content string) bool {
	fmt.Println("political filter")
	return strings.Contains(content, "习大大")
}