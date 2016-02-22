// Provides a convenient way to parse eagle schematic files.
package eagle

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type eagle struct {
	Schematic 	Schematic 	`xml:"drawing>schematic"`
	Version 	string 		`xml:"version,attr"`
}


// ----------------------------------------------------------------------------------
//  constructors
// ----------------------------------------------------------------------------------

func Load(file string) (*Schematic, error) {
	// open the xml file for parsing
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// read whole file in order to parse it into an object
	fileContent, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	// parse into object
	var xmlDoc eagle
	err = xml.Unmarshal(fileContent, &xmlDoc)
	if err != nil {
		return nil, err
	}

	// each part needs a reference to the schematic
	for i,_ := range xmlDoc.Schematic.Parts {
		xmlDoc.Schematic.Parts[i].schematic = &xmlDoc.Schematic
	}

	return &xmlDoc.Schematic, nil
}
