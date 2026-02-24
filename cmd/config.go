package cmd

import (
	"fmt"

	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/messages"
)

func getConfig() config.Configs {
	var cfg config.Configs
	cfg, err := config.LoadUserConfig()
	switch {
	case err != nil:
		messages.ExitErr(fmt.Errorf("unable to load user config: %w", err))
	case cfg != nil:
		cfg, err = cfg.WithDefaults()
		if err != nil {
			messages.ExitErr(fmt.Errorf("unable to merge user config: %w", err))
		}
	default:
		return config.Defaults()
	}
	return cfg
}

