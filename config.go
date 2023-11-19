package gopublicfield

import "strings"

type config struct {
	pkgs     stringsFlag
	onlyPkgs bool
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

	for _, pgk := range s {
		pgk = strings.TrimSpace(pgk)
		pgk = strings.TrimSuffix(pgk, "/") + "/"

		res = append(res, pgk)
	}

	return res
}

const (
	pkgsDesc     = "List of packages, which should use publicfield to initiate struct."
	onlyPkgsDesc = "Only pkg packages should use publicfield to initiate struct."
)
