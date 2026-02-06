package config

import (
	"embed"

	"github.com/outscale/gli/pkg/errors"
	"go.yaml.in/yaml/v4"
)

//go:generate go run generate/oapi/main.go generate/oapi/defaults.yaml defaults_oapi.yaml
//go:embed defaults_oapi.yaml
var f embed.FS

func Defaults() Configs {
	defaults := Configs{}
	for _, provider := range []string{"oapi"} {
		data, _ := f.ReadFile("defaults_" + provider + ".yaml")
		var cfg Config
		err := yaml.Unmarshal(data, &cfg)
		if err != nil {
			errors.ExitErr(err)
		}
		defaults[provider] = cfg
	}
	return defaults
}
