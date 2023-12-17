package app

import (
	"publicfield/packageGlobs/nested"
	"publicfield/packageGlobs/pkg"
	"publicfield/packageGlobs/pkg/blocked_nested"
)

func main() {
	nBlocked := pkg.Struct{}
	nBlocked.Int++ // want `Field 'Int' in pkg.Struct can be changed only inside nested package.`

	nBlockedPtr := &pkg.Struct{}
	nBlockedPtr.Int++ // want `Field 'Int' in pkg.Struct can be changed only inside nested package.`

	nBlockedNested := blocked_nested.Struct{}
	nBlockedNested.Int += 1 // want `Field 'Int' in blocked_nested.Struct can be changed only inside nested package.`

	n := nested.Struct{}
	n.Int-- // want `Field 'Int' in nested.Struct can be changed only inside nested package.`
}
