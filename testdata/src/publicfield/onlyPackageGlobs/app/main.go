package app

import (
	"publicfield/onlyPackageGlobs/nested"
	"publicfield/onlyPackageGlobs/pkg"
)

func main() {
	nBlocked := pkg.Struct{}
	nBlocked.Int++ // want `Field 'Int' in pkg.Struct can be changes only inside nested package.`

	n := nested.Struct{}
	n.Int--
}
