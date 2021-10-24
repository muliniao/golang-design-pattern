package standard

type ITemplate interface {
	Fun1()
	Fun2()
	Result()
}

// TemplateFather 父类模板,实现Result()
// Fun1()和Fun2()
type TemplateFather struct {
	template ITemplate
}

func NewTemplateFather(template ITemplate) *TemplateFather {
	return &TemplateFather{
		template: template,
	}
}

func (t *TemplateFather) Result() {
	t.template.Fun1()
	t.template.Fun2()
}
