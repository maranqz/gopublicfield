package app

import (
	"publicfield/onlyPkgs/nested"
	"publicfield/onlyPkgs/pkg"
)

func main() {
	nBlocked := pkg.Struct{}
	nBlocked.Int++ // want `Field 'Int' in pkg.Struct can be changes only inside nested package.`

	n := nested.Struct{}
	n.Int--
}
