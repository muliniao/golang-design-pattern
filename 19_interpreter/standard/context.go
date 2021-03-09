package standard

type Context struct {
	expression AbstractExpression
}

/**
数据初始化
*/
func NewContext() {

}

/**
调用相关表达式类的解释方法
*/
func (context *Context) Operation(info string) {

}
