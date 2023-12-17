package gopublicfield

import (
	"fmt"
	"strings"

	"github.com/gobwas/glob"
)

type globsFlag struct {
	globsString []string
	globs       []glob.Glob
}

func (g globsFlag) String() string {
	return strings.Join(g.globsString, ", ")
}

func (g *globsFlag) Set(globString string) error {
	globString = strings.TrimSpace(globString)

	compiled, err := glob.Compile(globString)
	if err != nil {
		return fmt.Errorf("unable to compile globs %s: %w", globString, err)
	}

	g.globsString = append(g.globsString, globString)
	g.globs = append(g.globs, compiled)

	return nil
}

func (g globsFlag) Value() []glob.Glob {
	return g.globs
}
