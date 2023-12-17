package nested

type Struct struct {
	Int    int
	IntPtr *int
}

func (s Struct) Increment() Struct {
	s.Int++

	return s
}
