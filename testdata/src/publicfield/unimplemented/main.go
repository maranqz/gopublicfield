package unimplemented

import "publicfield/unimplemented/nested"

func ByPtr() {
	i := 1
	st := nested.Struct{
		IntPtr: &i,
	}

	stInt := &st.Int // lint alert can be here
	*stInt = 2       // want `Field 'Int' in nested.Struct can be changed only inside nested package.`

	iPtr := st.IntPtr
	*iPtr = 2 // want `Field 'Int' in nested.Struct can be changed only inside nested package.`
}
