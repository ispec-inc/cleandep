package cleandep

import (
	"os"
	"strconv"

	"github.com/goccy/go-yaml"
	"github.com/mattn/go-zglob"
	"golang.org/x/tools/go/analysis"
)

const doc = "cleandep detects illegal dependencies"

var Analyzer = &analysis.Analyzer{
	Name:     "cleandep",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

func decodeConfig(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	cfg, err := decodeConfig(".cleandep.yaml")
	if err != nil {
		return nil, err
	}

	rules := make(map[string][]string, len(cfg.Rules))
	for _, rule := range cfg.Rules {
		rules[rule.Source] = rule.IllegalDependencies
	}

	src := pass.Pkg.Path()

	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			dst, err := strconv.Unquote(imp.Path.Value)
			if err != nil {
				continue
			}

			for srcPattern, dstPatterns := range rules {
				srcMatched, err := zglob.Match(srcPattern, src)
				if err != nil {
					continue
				}
				if !srcMatched {
					continue
				}
				for _, dstPattern := range dstPatterns {
					dstMatched, err := zglob.Match(dstPattern, dst)
					if err != nil {
						continue
					}
					if dstMatched {
						pass.Reportf(imp.Pos(), "package '%s' cannot depend on package '%s'", src, dst)
					}
				}
			}
		}
	}

	return nil, nil
}
