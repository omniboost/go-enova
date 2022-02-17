package enova

import (
	"encoding/xml"
)

type RequestEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"s:Header"`
	Body   Body   `xml:"s:Body"`
}

func NewRequestEnvelope() RequestEnvelope {
	return RequestEnvelope{
		Header: NewHeader(),
	}
}

type ResponseEnvelope struct {
	XMLName xml.Name

	Header Header `xml:"Header"`
	Body   Body   `xml:"Body"`
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

type Header struct {
	Action Action `xml:"Action"`
}

type Action string

func (act Action) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.microsoft.com/ws/2005/05/addressing/none"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s:mustUnderstand"}, Value: "1"})

	type alias Action
	a := alias(act)
	return e.EncodeElement(a, start)
}

func NewHeader() Header {
	return Header{
		Action: Action("http://tempuri.org/IMethodInvokerService/InvokeServiceMethod"),
	}
}

type ActionBody interface{}
