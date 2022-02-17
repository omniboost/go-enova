package enova

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-enova/utils"
)

func (c *Client) NewListlaendercodes() Listlaendercodes {
	r := Listlaendercodes{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type Listlaendercodes struct {
	client      *Client
	queryParams *ListlaendercodesQueryParams
	pathParams  *ListlaendercodesPathParams
	method      string
	headers     http.Header
	requestBody ListlaendercodesBody
}

func (r Listlaendercodes) NewQueryParams() *ListlaendercodesQueryParams {
	return &ListlaendercodesQueryParams{}
}

type ListlaendercodesQueryParams struct {
}

func (p ListlaendercodesQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *Listlaendercodes) QueryParams() QueryParams {
	return r.queryParams
}

func (r Listlaendercodes) NewPathParams() *ListlaendercodesPathParams {
	return &ListlaendercodesPathParams{}
}

type ListlaendercodesPathParams struct {
}

func (p *ListlaendercodesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *Listlaendercodes) PathParams() PathParams {
	return r.pathParams
}

func (r *Listlaendercodes) SetMethod(method string) {
	r.method = method
}

func (r *Listlaendercodes) Method() string {
	return r.method
}

func (r Listlaendercodes) NewRequestBody() ListlaendercodesBody {
	return ListlaendercodesBody{}
}

type ListlaendercodesBody struct {
	XMLName xml.Name `xml:"car:listlaendercodes"`
	Mandant string   `xml:"car:mandant"`
}

func (r *Listlaendercodes) RequestBody() *ListlaendercodesBody {
	return &r.requestBody
}

func (r *Listlaendercodes) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *Listlaendercodes) SetRequestBody(body ListlaendercodesBody) {
	r.requestBody = body
}

func (r *Listlaendercodes) NewResponseBody() *ListlaendercodesResponseBody {
	return &ListlaendercodesResponseBody{}
}

type ListlaendercodesResponseBody struct {
	XMLName xml.Name `xml:"listlaendercodesResponse"`

	// ListlaendercodesResult struct {
	// 	Laender []struct {
	// 		Name string `xml:"name,attr"`
	// 	} `xml:"laender>country"`
	// } `xml:"listlaendercodesResult"`

	Laender []struct {
		Name    string `xml:"name,attr"`
		IntName string `xml:"IntName,attr"`
		Capital string `xml:"Capital,attr"`
		ISO2    string `xml:"ISO-2,attr"`
		ISO3    string `xml:"ISO-3,attr"`
		KFZ     string `xml:"KFZ,attr"`
	} `xml:"listlaendercodesResult>laender>country"`
}

func (r *Listlaendercodes) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *Listlaendercodes) Do() (ListlaendercodesResponseBody, error) {
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
	return *responseBody, err
}
