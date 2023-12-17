package nested

import (
	"publicfield/packageGlobsOnly/pkg"
)

type Struct struct {
	Int int
}

func callNested1() {
	n := pkg.Struct{}
	n.Int = 1 // want `Field 'Int' in pkg.Struct can be changed only inside nested package.`
}
