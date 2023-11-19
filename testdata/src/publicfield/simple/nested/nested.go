package nested

type Struct struct {
	Int   int
	Slice []int
	Map   map[int]struct{}
}

func (s Struct) Increment() Struct {
	s.Int++

	return s
}
