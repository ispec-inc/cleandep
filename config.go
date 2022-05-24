package cleandep

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Source              string   `yaml:"source"`
	IllegalDependencies []string `yaml:"illegal_dependencies"`
}
