package meldeschein

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-avs-meldeschein/utils"
)

func (c *Client) NewGetMeldescheine() GetMeldescheine {
	r := GetMeldescheine{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetMeldescheine struct {
	client      *Client
	queryParams *GetMeldescheineQueryParams
	pathParams  *GetMeldescheinePathParams
	method      string
	headers     http.Header
	requestBody GetMeldescheineBody
}

func (r GetMeldescheine) NewQueryParams() *GetMeldescheineQueryParams {
	return &GetMeldescheineQueryParams{}
}

type GetMeldescheineQueryParams struct {
}

func (p GetMeldescheineQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetMeldescheine) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetMeldescheine) NewPathParams() *GetMeldescheinePathParams {
	return &GetMeldescheinePathParams{}
}

type GetMeldescheinePathParams struct {
}

func (p *GetMeldescheinePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetMeldescheine) PathParams() PathParams {
	return r.pathParams
}

func (r *GetMeldescheine) SetMethod(method string) {
	r.method = method
}

func (r *GetMeldescheine) Method() string {
	return r.method
}

func (r GetMeldescheine) NewRequestBody() GetMeldescheineBody {
	return GetMeldescheineBody{
		Identifikation: Identifikation{
			Erzeugung:     Date{time.Now()},
			Kurverwaltung: 0,
			BenutzerID:    1,
			Verarbeitung:  "MS-HOLEN",
		},
	}
}

type GetMeldescheineBody struct {
	XMLName        xml.Name       `xml:"ns:holeMeldescheineRequest"`
	Identifikation Identifikation `xml:"identifikation"`
	AnfrageDaten   struct {
		// You have a CHOICE of the next 3 items at this level
		Buchungsnummer    string `xml:"buchungsnummer,omitempty"`
		Meldescheinnummer string `xml:"meldescheinnummer,omitempty"`
		OrtID             string `xml:"ort-id,omitempty"`
	} `xml:"anfragedaten"`
}

func (r *GetMeldescheine) RequestBody() *GetMeldescheineBody {
	return &r.requestBody
}

func (r *GetMeldescheine) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetMeldescheine) SetRequestBody(body GetMeldescheineBody) {
	r.requestBody = body
}

func (r *GetMeldescheine) NewResponseBody() *GetMeldescheineResponseBody {
	return &GetMeldescheineResponseBody{}
}

type GetMeldescheineResponseBody struct {
	XMLName xml.Name `xml:"holeMeldescheineResponse"`

	Identifikation  Identifikation  `xml:"identifikation"`
	Fehlermeldungen Fehlermeldungen `xml:"fehlermeldungen"`
}

func (r *GetMeldescheine) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetMeldescheine) Do() (GetMeldescheineResponseBody, error) {
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
