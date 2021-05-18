package cardxperts

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-cardxperts/utils"
	"github.com/pkg/errors"
)

func (c *Client) NewMeldedatenRequest() MeldedatenRequest {
	r := MeldedatenRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldedatenRequest struct {
	client      *Client
	queryParams *MeldedatenRequestQueryParams
	pathParams  *MeldedatenRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldedatenRequestBody
}

func (r MeldedatenRequest) NewQueryParams() *MeldedatenRequestQueryParams {
	return &MeldedatenRequestQueryParams{}
}

type MeldedatenRequestQueryParams struct {
}

func (p MeldedatenRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *MeldedatenRequest) QueryParams() QueryParams {
	return r.queryParams
}

func (r MeldedatenRequest) NewPathParams() *MeldedatenRequestPathParams {
	return &MeldedatenRequestPathParams{}
}

type MeldedatenRequestPathParams struct {
}

func (p *MeldedatenRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MeldedatenRequest) PathParams() PathParams {
	return r.pathParams
}

func (r *MeldedatenRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldedatenRequest) Method() string {
	return r.method
}

func (r MeldedatenRequest) NewRequestBody() MeldedatenRequestBody {
	return MeldedatenRequestBody{
		Meldedaten: Meldedaten{
			Version: "5",
		},
	}
}

type MeldedatenRequestBody struct {
	XMLName          xml.Name   `xml:"car:Meldedaten"`
	Meldedaten       Meldedaten `xml:"-"`
	MeldedatenString string     `xml:"car:meldedatenstring"`
}

func (mrb MeldedatenRequestBody) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	b, err := xml.MarshalIndent(mrb.Meldedaten, "", "  ")
	if err != nil {
		return errors.WithStack(err)
	}
	mrb.MeldedatenString = string(b)

	type alias struct {
		XMLName          xml.Name `xml:"car:Meldedaten"`
		MeldedatenString CData    `xml:"car:meldedatenstring"`
	}
	a := alias{
		MeldedatenString: CData(b),
	}
	return e.Encode(a)
}

func (r *MeldedatenRequest) RequestBody() *MeldedatenRequestBody {
	return &r.requestBody
}

func (r *MeldedatenRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *MeldedatenRequest) SetRequestBody(body MeldedatenRequestBody) {
	r.requestBody = body
}

func (r *MeldedatenRequest) NewResponseBody() *MeldedatenRequestResponseBody {
	return &MeldedatenRequestResponseBody{}
}

type MeldedatenRequestResponseBody struct {
	XMLName xml.Name `xml:"MeldedatenResponse"`

	Error struct {
		Fehlertext string `xml:"Fehlertext,attr"`
	} `xml:"MeldedatenResult>error"`
	NewDataSet struct {
		Return struct {
			Number      int    `xml:"Number"`
			Description string `xml:"Description"`
		}
	} `xml:"MeldedatenResult>NewDataSet"`
}

func (r *MeldedatenRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *MeldedatenRequest) Do() (MeldedatenRequestResponseBody, error) {
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

	if responseBody.Error.Fehlertext != "" {
		return *responseBody, errors.New(responseBody.Error.Fehlertext)
	}

	return *responseBody, nil
}
