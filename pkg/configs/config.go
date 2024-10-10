package configs

import (
	"github.com/KoNekoD/go-deptrac/pkg/hooks"
)

type config struct {
	ConfigFileHook
	hooks.TemplateFileHook
}

var Instance *config

func init() {
	Instance = &config{
		ConfigFileHook:   NewConfigFileHook(),
		TemplateFileHook: hooks.NewTemplateFileHook(),
	}
}
