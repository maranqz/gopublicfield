package gopublicfield_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/maranqz/gopublicfield"
)

func TestLinterSuite(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()

	tests := map[string]struct {
		pkgs    []string
		prepare func(t *testing.T, a *analysis.Analyzer)
	}{
		"simple": {pkgs: []string{"simple/..."}},
		"pgks": {
			pkgs: []string{"pgks/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) {
				if err := a.Flags.Set("pkgs", "publicfield/pgks/pkg"); err != nil {
					t.Fatal(err)
				}
			},
		},
		"onlyPkgs": {
			pkgs: []string{"onlyPkgs/app/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) {
				if err := a.Flags.Set("p", "publicfield/onlyPkgs/pkg"); err != nil {
					t.Fatal(err)
				}

				if err := a.Flags.Set("op", "true"); err != nil {
					t.Fatal(err)
				}

				if err := a.Flags.Set("onlyPkgs", "true"); err != nil {
					t.Fatal(err)
				}
			},
		},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dirs := make([]string, 0, len(tt.pkgs))

			for _, pkg := range tt.pkgs {
				dirs = append(dirs, filepath.Join(testdata, "src", "publicfield", pkg))
			}

			analyzer := gopublicfield.NewAnalyzer()

			if tt.prepare != nil {
				tt.prepare(t, analyzer)
			}

			analysistest.Run(t, TestdataDir(),
				analyzer, dirs...)
		})
	}
}

func TestdataDir() string {
	_, testFilename, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get current test filename")
	}

	return filepath.Join(filepath.Dir(testFilename), "testdata")
}
