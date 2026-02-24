package config

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"dario.cat/mergo"
	"github.com/goccy/go-yaml"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
)

var CanConfigure = true

func LoadUserConfig() (Configs, error) {
	cfg := Configs{}
	root, err := os.UserConfigDir()
	if err != nil {
		CanConfigure = false
		debug.Println("Unable to compute user config dir", err)
		return nil, nil
	}
	path := filepath.Join(root, "osc", "config.yaml")
	data, err := os.ReadFile(path) //nolint:gosec
	if errors.Is(err, fs.ErrNotExist) {
		debug.Println("no user config file found", path)
		return nil, nil
	}
	debug.Println("loading user config from", path)
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		messages.ExitErr(err)
	}
	return cfg, nil
}

func (cfg Configs) WithDefaults() (Configs, error) {
	def := Defaults()
	for provider := range def {
		pcfg := cfg[provider]
		err := mergo.Merge(&pcfg, def[provider])
		if err != nil {
			return nil, err
		}
		cfg[provider] = pcfg
	}
	return cfg, nil
}

func (cfg Configs) For(provider string) Config {
	return cfg[provider]
}
