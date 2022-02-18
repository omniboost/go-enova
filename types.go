package enova

import (
	"encoding/xml"
)

type UpdateParamsData struct {
	XMLName xml.Name `xml:"UpdateParamsData"`

	Rows Rows `xml:"Rows>Row"`
}

func (d UpdateParamsData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias UpdateParamsData
	a := alias(d)
	b, err := xml.Marshal(a)
	if err != nil {
		return err
	}

	return e.EncodeElement(b, start)
}

type Rows []Row

type Row struct {
	XML interface{} `xml:"Xml"`
}
