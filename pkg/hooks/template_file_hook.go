package hooks

type templateFileHook struct{}
type TemplateFileHook interface {
	GetTemplateFile() string
}

func NewTemplateFileHook() TemplateFileHook {
	return &templateFileHook{}
}

const DefaultTemplateFile = "deptrac_template.yaml"

func (h *templateFileHook) GetTemplateFile() string {

	return DefaultTemplateFile
}
