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
		prepare func(t *testing.T, a *analysis.Analyzer) error
	}{
		"simple": {pkgs: []string{"simple/..."}},
		"packageGlobs": {
			pkgs: []string{"packageGlobs/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) error {
				return a.Flags.Set("packageGlobs", "publicfield/packageGlobs/pkg/**")
			},
		},
		"packageGlobsOnly": {
			pkgs: []string{"packageGlobsOnly/app/..."},
			prepare: func(t *testing.T, a *analysis.Analyzer) error {
				if err := a.Flags.Set("packageGlobs", "publicfield/packageGlobsOnly/pkg/**"); err != nil {
					return err
				}

				return a.Flags.Set("packageGlobsOnly", "true")
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
				err := tt.prepare(t, analyzer)
				if err != nil {
					t.Fatal(err)
				}
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
