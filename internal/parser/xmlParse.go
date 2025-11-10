// File: internal/parser/xml_parser.go
package parser

import (
	"encoding/xml"
	"log"
	// "fmt"
	"io"

	"github.com/AbhaySingh002/Pollen/internal/types"
)



func ParseStreamFromReader(r io.Reader) {
	decoder := xml.NewDecoder(r)

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}else if err != nil {
			panic(err)
		}
		switch se:= tok.(type){
		case xml.StartElement:
			if se.Name.Local == "Projectname"{
				var Title string
				if err := decoder.DecodeElement(&Title, &se); err != nil {
					panic(err)
				}
				log.Printf("Project Title : %s\n",Title)
			}
			if se.Name.Local == "Step" {
				var step types.Step
				if err := decoder.DecodeElement(&step, &se); err != nil {
					panic(err)
				}
				switch step.Type {
				case "file":
					log.Printf("Writing the file : %s\n", step.Path)
				case "command":
					log.Printf("Executing the command : %s\n", step.Command)
				}
			}
		}

	}
    
}