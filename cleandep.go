package cleandep

import (
	"go/ast"
	"os"
	"strconv"

	"github.com/goccy/go-yaml"
	"github.com/mattn/go-zglob"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "cleandep detects illegal dependencies"

var Analyzer = &analysis.Analyzer{
	Name: "cleandep",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
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
		rules[rule.Source] = rule.IllegalDestinations
	}

	src := pass.Pkg.Path()

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ImportSpec:
			dst, err := strconv.Unquote(n.Path.Value)
			if err != nil {
				return
			}

			for srcPattern, dstPatterns := range rules {
				srcMatched, err := zglob.Match(srcPattern, src)
				if err != nil {
					return
				}
				if !srcMatched {
					continue
				}
				for _, dstPattern := range dstPatterns {
					dstMatched, err := zglob.Match(dstPattern, dst)
					if err != nil {
						return
					}
					if dstMatched {
						pass.Reportf(n.Pos(), "package '%s' cannot depend on package '%s'", src, dst)
					}
				}
			}
		}
	})

	return nil, nil
}
