package gopublicfield

import "strings"

type config struct {
	packageGlobs     stringsFlag
	onlyPackageGlobs bool
}

type stringsFlag []string

func (s stringsFlag) String() string {
	return strings.Join(s, ", ")
}

func (s *stringsFlag) Set(value string) error {
	*s = append(*s, value)

	return nil
}

func (s stringsFlag) Value() []string {
	res := make([]string, 0, len(s))

	for _, str := range s {
		res = append(res, strings.TrimSpace(str))
	}

	return res
}

const (
	packageGlobsDesc     = "List of glob packages, in which public fields can be written by other packages in the same glob pattern."
	onlyPackageGlobsDesc = "Only public fields in glob packages cannot be written by other packages."
)
