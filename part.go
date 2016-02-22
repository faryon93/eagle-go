package eagle


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Part struct {
	Name 		string `xml:"name,attr"`
	Library 	string `xml:"library,attr"`
	Deviceset 	string `xml:"deviceset,attr"`

	schematic	*Schematic
}


// ----------------------------------------------------------------------------------
//  informational members
// ----------------------------------------------------------------------------------

func (p *Part) getLibrary() *Library {
	if (p.schematic == nil) {
		return nil
	}

	for _, l := range p.schematic.Libraries {
		if l.Name == p.Library {
			return &l
		}
	}

	return nil
}

func (p *Part) GetDeviceset() *Deviceset {
	return p.getLibrary().GetDeviceset(p.Deviceset)
}