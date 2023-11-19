package local

type Struct struct {
	Int int
}

func NewStructInline() (Struct, int, error) {
	return Struct{
		Int: 1,
	}, 10, nil
}

func NewStruct() (Struct, int, error) {
	res := Struct{}

	res.Int = 1

	return res, 0, nil
}

func (s *Struct) UpdatePtr() {
	s.Int = 1

}

func (s Struct) Update() {
	s.Int = 1
}

func UpdatePtr(s *Struct) {
	s.Int = 1 // want `Field 'Int' in local.Struct can be changes only inside Factory or Struct methods.`
}

func Update(s Struct) Struct {
	s.Int = 1 // want `Field 'Int' in local.Struct can be changes only inside Factory or Struct methods.`

	return s
}
