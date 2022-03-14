package enova

import (
	"encoding/xml"
)

type UpdateParamsData struct {
	XMLName xml.Name `xml:"UpdateParamsData"`

	Rows Rows `xml:"Rows>Row>Xml>XMLName"`
}

func (d UpdateParamsData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias UpdateParamsData
	a := alias(d)
	b, err := xml.Marshal(a)
	if err != nil {
		return err
	}

	return e.EncodeElement(string(b), start)
}

type Rows []Row

type Row interface{}
