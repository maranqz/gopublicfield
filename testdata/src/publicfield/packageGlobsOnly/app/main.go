package app

import (
	"publicfield/packageGlobsOnly/nested"
	"publicfield/packageGlobsOnly/pkg"
)

func main() {
	nBlocked := pkg.Struct{}
	nBlocked.Int++ // want `Field 'Int' in pkg.Struct can be changed only inside nested package.`

	n := nested.Struct{}
	n.Int--
}
