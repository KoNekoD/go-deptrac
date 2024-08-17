package config

type config struct {
	ConfigFileHook
	TemplateFileHook
}

var Instance *config

func init() {
	Instance = &config{
		ConfigFileHook:   NewConfigFileHook(),
		TemplateFileHook: NewTemplateFileHook(),
	}
}
