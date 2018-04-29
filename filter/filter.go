package filter

/*
Filter ...
*/
type Filter struct {
	HasPhoto           bool
	InContact          bool
	Favouraite         bool
	CompatibilityScore float32
	Age                int32
	Height             int32
	Distance           int32
}

/*
NewFilter ...
*/
func NewFilter(hasPhoto, inContact, favouraite bool, score float32, age, height, distance int32) *Filter {
	return &Filter{
		HasPhoto: hasPhoto,
		InContact: inContact,
		Favouraite: favouraite,
		CompatibilityScore: score,
		Age: age,
		Height: height,
		Distance: distance,
	}
}
