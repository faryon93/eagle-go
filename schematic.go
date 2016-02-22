package eagle


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Schematic struct {
	Parts		[]Part 		`xml:"parts>part"`
	Nets 		[]Net 		`xml:"sheets>sheet>nets>net"`
	Libraries 	[]Library 	`xml:"libraries>library"`
}


// ----------------------------------------------------------------------------------
//  informational member functions
// ----------------------------------------------------------------------------------

func (s *Schematic) GetPart(name string) *Part {
	for _, p := range s.Parts {
		if p.Name == name {
			return &p
		}
	}

	return nil
}