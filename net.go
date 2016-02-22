package eagle

// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Net struct {
	Name 		string 		`xml:"name,attr"`
	PinRefs 	[]PinRef 	`xml:"segment>pinref"`
}

type PinRef struct {
	Part 	string	`xml:"part,attr"`
	Pin 	string	`xml:"pin,attr"`
}


// ----------------------------------------------------------------------------------
//  informational members
// ----------------------------------------------------------------------------------

func (n *Net) GetPartPin(part string) string {
	for _, pinRef := range n.PinRefs {
		if (pinRef.Part == part) {
			return pinRef.Pin
		}

	}

	return ""
}