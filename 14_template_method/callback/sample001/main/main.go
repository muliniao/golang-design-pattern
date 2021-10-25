package main


// 小明求助小红帮忙计算两数之和
// 计算后，回调回来。填写结果
func main() {
	a, b := 168, 291
	s := NewStudent("小明")
	s.CallHelp(a, b)
}
