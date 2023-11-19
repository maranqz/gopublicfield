package pkg

import "publicfield/packageGlobs/pkg/blocked_nested"

type Struct struct {
	Int int
}

func CallNested2() {
	n := blocked_nested.Struct{}
	n.Int = 1
}
