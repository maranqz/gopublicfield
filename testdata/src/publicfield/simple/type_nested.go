package simple

import "publicfield/simple/nested"

type AliasStruct = nested.Struct

func typeNested() {
	as := AliasStruct{}
	as.Int++ // want `Field 'Int' in nested.Struct can be changed only inside nested package.`
}
