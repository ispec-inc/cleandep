package main

import (
	"github.com/ispec-inc/cleandep"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(cleandep.Analyzer) }
