package cleandep_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/ispec-inc/cleandep"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, cleandep.Analyzer, "a", "a/b")
}
