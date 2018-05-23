package filter

/*
Filter ...
*/
type Filter struct {
	HasPhoto           bool
	InContact          bool
	Favouraite         bool
	CompatibilityScore float32
	MinAge             int32
	MaxAge             int32
	Height             int32
	Distance           int32
}

/*
NewFilter ...
*/
func NewFilter(hasPhoto, inContact, favouraite bool, score float32, minAge, maxAge, height, distance int32) *Filter {
	return &Filter{
		HasPhoto:           hasPhoto,
		InContact:          inContact,
		Favouraite:         favouraite,
		CompatibilityScore: score,
		MinAge:             minAge,
		MaxAge:             maxAge,
		Height:             height,
		Distance:           distance,
	}
}
