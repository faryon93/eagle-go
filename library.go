package eagle

// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Library struct {
	Name 			string 			`xml:"name,attr"`
	Description 	string 			`xml:"description"`
	Devicesets 		[]Deviceset 	`xml:"devicesets>deviceset"`
}


// ----------------------------------------------------------------------------------
//  informational members
// ----------------------------------------------------------------------------------

func (l *Library) GetDeviceset(name string) *Deviceset {
	for _, d := range l.Devicesets {
		if (d.Name == name) {
			return &d
		}
	}

	return nil
}