package cleandep

type Config struct {
	Rules []Rule `yaml:"rules"`
}

type Rule struct {
	Source              string   `yaml:"source"`
	IllegalDestinations []string `yaml:"illegal_destinations"`
}
