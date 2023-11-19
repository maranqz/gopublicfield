package simple

import "publicfield/simple/nested"

type Struct struct {
	Field int
}

func local() {
	n := Struct{}
	n.Field = 1
}

func NestedInt() {
	n := nested.Struct{
		Int: 1001,
	}

	n = n.Increment()

	n.Int = 1 // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int,    // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
		n.Int = // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
		1, 2
	n.Int++     // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int--     // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int += 1  // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int -= 1  // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int &= 1  // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int ^= 1  // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int |= 1  // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int >>= 1 // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
	n.Int <<= 1 // want `Field 'Int' in nested.Struct can be changes only inside nested package.`

}

func NestedIntPtr() {
	nPtr1 := &nested.Struct{}
	nPtr1.Int-- // want `Field 'Int' in nested.Struct can be changes only inside nested package.`

	n2 := nested.Struct{}

	nPtr2 := &n2
	nPtr22 := nPtr2
	nPtr22.Int++ // want `Field 'Int' in nested.Struct can be changes only inside nested package.`

	nPtr2_2 := &nPtr2
	nPtr2_2_2 := &nPtr2_2
	(***nPtr2_2_2).Int++ // want `Field 'Int' in nested.Struct can be changes only inside nested package.`
}

func NestedSlice() {
	n := nested.Struct{}
	n.Slice = append(n.Slice, 1) // want `Field 'Slice' in nested.Struct can be changes only inside nested package.`

	// Skipped
	n.Slice[0] = 1
}

func NestedMap() {
	n := nested.Struct{}
	n.Map = map[int]struct{}{ // want `Field 'Map' in nested.Struct can be changes only inside nested package.`
		1: {},
	}

	// Skipped
	n.Map[1] = struct{}{}
}
