package sample

type SensitiveWordFilter interface {
	doFilter(content string) bool
}
