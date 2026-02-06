package config

type Alias struct {
	Group   string   `yaml:"group"`
	Use     string   `yaml:"use"`
	Short   string   `yaml:"short"`
	Command []string `yaml:"command"`
}

type FlagConfig struct {
	AppliesTo  string `yaml:"applies_to"`
	TrimPrefix string `yaml:"trim_prefix"`
}

type Config struct {
	Aliases []Alias      `yaml:"aliases"`
	Flags   []FlagConfig `yaml:"flags"`
}

type Configs map[string]Config

func For(provider string) Config {
	return Defaults()[provider]
}
