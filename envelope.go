package enova

import (
	"encoding/xml"
)

type RequestEnvelope struct {
	XMLName xml.Name

	// Header Header `xml:"s:Header"`
	Body Body `xml:"s:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		// Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	// Header Header `xml:"Header"`
	Body Body `xml:"Body"`
}

func (env RequestEnvelope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "s:Envelope"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:s"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias RequestEnvelope
	a := alias(env)
	return e.EncodeElement(a, start)
}

type Body struct {
	ActionBody interface{} `xml:",any"`
}

// type Header struct {
// }

// func NewHeader() Header {
// 	return Header{}
// }

type ActionBody interface{}
