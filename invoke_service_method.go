package enova

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-enova/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewInvokeServiceMethodRequest() InvokeServiceMethodRequest {
	r := InvokeServiceMethodRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvokeServiceMethodRequest struct {
	client      *Client
	queryParams *InvokeServiceMethodRequestQueryParams
	pathParams  *InvokeServiceMethodRequestPathParams
	method      string
	headers     http.Header
	requestBody InvokeServiceMethodRequestBody
}

func (r InvokeServiceMethodRequest) NewQueryParams() *InvokeServiceMethodRequestQueryParams {
	return &InvokeServiceMethodRequestQueryParams{}
}

type InvokeServiceMethodRequestQueryParams struct {
}

func (p InvokeServiceMethodRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvokeServiceMethodRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r InvokeServiceMethodRequest) NewPathParams() *InvokeServiceMethodRequestPathParams {
	return &InvokeServiceMethodRequestPathParams{}
}

type InvokeServiceMethodRequestPathParams struct {
}

func (p *InvokeServiceMethodRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvokeServiceMethodRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *InvokeServiceMethodRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvokeServiceMethodRequest) Method() string {
	return r.method
}

func (r InvokeServiceMethodRequest) NewRequestBody() InvokeServiceMethodRequestBody {
	return InvokeServiceMethodRequestBody{}
}

type InvokeServiceMethodRequestBody struct {
	XMLName xml.Name `xml:"InvokeServiceMethod"`

	InvokerParams InvokerParams `xml:"invokerParams"`
}

func (rb InvokeServiceMethodRequestBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "InvokeServiceMethod"}

	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://tempuri.org/"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias InvokeServiceMethodRequestBody
	a := alias(rb)
	return e.EncodeElement(a, start)
}

type InvokerParams struct {
	// ClientAddress struct {
	// 	Xmlns string `xml:"xmlns,attr"`
	// 	Nil   string `xml:"nil,attr"`
	// } `xml:"ClientAddress"`
	// ContextHandle struct {
	// 	Xmlns string `xml:"xmlns,attr"`
	// 	Nil   string `xml:"nil,attr"`
	// } `xml:"ContextHandle"`
	DatabaseHandle DatabaseHandle `xml:"DatabaseHandle"`
	// IpFilters struct {
	// 	Xmlns string `xml:"xmlns,attr"`
	// 	Nil   string `xml:"nil,attr"`
	// } `xml:"IpFilters"`
	MethodArgs MethodArgs `xml:"MethodArgs"`
	MethodName MethodName `xml:"MethodName"`
	Operator   Operator   `xml:"Operator"`
	Password   Password   `xml:"Password"`
	// ServiceName struct {
	// 	Xmlns string `xml:"xmlns,attr"`
	// } `xml:"ServiceName"`
	// ConnectionInfo struct {
	// 	Nil string `xml:"nil,attr"`
	// } `xml:"ConnectionInfo"`
}

func (s InvokerParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:d4p1"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Net.Types"},
		{Name: xml.Name{Space: "", Local: "xmlns:i"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias InvokerParams
	a := alias(s)
	return e.EncodeElement(a, start)
}

type DatabaseHandle string

func (s DatabaseHandle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias DatabaseHandle
	a := alias(s)
	return e.EncodeElement(a, start)
}

type MethodArgs map[string]interface{}

func (s MethodArgs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:d5p1"}, Value: "http://schemas.microsoft.com/2003/10/Serialization/Arrays"},
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	list := struct {
		XMLName                 xml.Name                  `xml:"MethodArgs"`
		KeyValueOfstringanyType []KeyValueOfstringanyType `xml:"d5p1:KeyValueOfstringanyType"`
	}{}

	for k, v := range s {
		list.KeyValueOfstringanyType = append(list.KeyValueOfstringanyType,
			KeyValueOfstringanyType{
				Key: k,
				Value: value{
					I: v,
				},
			})
	}
	return e.EncodeElement(list, start)
}

type KeyValueOfstringanyType struct {
	Key   string `xml:"d5p1:Key"`
	Value value  `xml:"d5p1:Value"`
}

type value struct {
	I interface{}
}

func (v value) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns:d7p1"}, Value: "http://www.w3.org/2001/XMLSchema"},
		{Name: xml.Name{Space: "", Local: "i:type"}, Value: "d7p1:string"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	return e.EncodeElement(v.I, start)
}

type MethodName string

func (s MethodName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias MethodName
	a := alias(s)
	return e.EncodeElement(a, start)
}

type Operator string

func (s Operator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias Operator
	a := alias(s)
	return e.EncodeElement(a, start)
}

type Password string

func (s Password) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias Password
	a := alias(s)
	return e.EncodeElement(a, start)
}

type ServiceName string

func (s ServiceName) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	namespaces := []xml.Attr{
		{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.datacontract.org/2004/07/Soneta.Types"},
	}
	for _, ns := range namespaces {
		start.Attr = append(start.Attr, ns)
	}

	type alias ServiceName
	a := alias(s)
	return e.EncodeElement(a, start)
}

func (r *InvokeServiceMethodRequest) RequestBody() *InvokeServiceMethodRequestBody {
	return &r.requestBody
}

func (r *InvokeServiceMethodRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InvokeServiceMethodRequest) SetRequestBody(body InvokeServiceMethodRequestBody) {
	r.requestBody = body
}

func (r *InvokeServiceMethodRequest) NewResponseBody() *InvokeServiceMethodRequestResponseBody {
	return &InvokeServiceMethodRequestResponseBody{}
}

type InvokeServiceMethodRequestResponseBody struct {
	XMLName xml.Name `xml:"InvokeServiceMethodResponse"`
}

func (r *InvokeServiceMethodRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *InvokeServiceMethodRequest) Do() (InvokeServiceMethodRequestResponseBody, error) {
	var err error

	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	if err != nil {
		return *responseBody, errors.WithStack(err)
	}

	return *responseBody, nil
}
