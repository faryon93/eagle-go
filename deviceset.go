package eagle

// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Deviceset struct {
	Name 		string 	  `xml:"name,attr"`
	Connects 	[]Connect `xml:"devices>device>connects>connect"`
}

type Connect struct {
	Pin 	string 	`xml:"pin,attr"`
	Pad 	string 	`xml:"pad,attr"`
}


// ----------------------------------------------------------------------------------
//  informational members
// ----------------------------------------------------------------------------------

func (d *Deviceset) GetPad(pin string) string {
	for _, c := range d.Connects {
		if (c.Pin == pin) {
			return c.Pad
		}
	}

	return ""
}