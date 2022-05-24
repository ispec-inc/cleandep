package cleandep

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Package             string   `yaml:"package"`
	IllegalDependencies []string `yaml:"illegal_dependencies"`
}

func decodeConfig(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
