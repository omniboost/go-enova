package meldeschein

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-cardxperts/utils"
)

func (c *Client) NewGetMeldeschein() GetMeldeschein {
	r := GetMeldeschein{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetMeldeschein struct {
	client      *Client
	queryParams *GetMeldescheinQueryParams
	pathParams  *GetMeldescheinPathParams
	method      string
	headers     http.Header
	requestBody GetMeldescheinRequestBody
}

func (r GetMeldeschein) NewQueryParams() *GetMeldescheinQueryParams {
	return &GetMeldescheinQueryParams{}
}

type GetMeldescheinQueryParams struct {
}

func (p GetMeldescheinQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetMeldeschein) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetMeldeschein) NewPathParams() *GetMeldescheinPathParams {
	return &GetMeldescheinPathParams{}
}

type GetMeldescheinPathParams struct {
}

func (p *GetMeldescheinPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetMeldeschein) PathParams() PathParams {
	return r.pathParams
}

func (r *GetMeldeschein) SetMethod(method string) {
	r.method = method
}

func (r *GetMeldeschein) Method() string {
	return r.method
}

func (r GetMeldeschein) NewRequestBody() GetMeldescheinRequestBody {
	return GetMeldescheinRequestBody{
		Identifikation: Identifikation{
			Erzeugung:     Date{time.Now()},
			Kurverwaltung: 0,
			BenutzerID:    1,
			Verarbeitung:  "MS-HOLEN",
		},
	}
}

type GetMeldescheinRequestBody struct {
	XMLName        xml.Name       `xml:"ns:holeMeldeschein"`
	Identifikation Identifikation `xml:"identifikation"`
	AnfrageDaten   struct {
		// You have a CHOICE of the next 3 items at this level
		Buchungsnummer    string `xml:"buchungsnummer,omitempty"`
		Meldescheinnummer string `xml:"meldescheinnummer,omitempty"`
		OrtID             int    `xml:"ort-id,omitempty"`
	} `xml:"anfragedaten"`
}

func (r *GetMeldeschein) RequestBody() *GetMeldescheinRequestBody {
	return &r.requestBody
}

func (r *GetMeldeschein) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetMeldeschein) SetRequestBody(body GetMeldescheinRequestBody) {
	r.requestBody = body
}

func (r *GetMeldeschein) NewResponseBody() *GetMeldescheinResponseBody {
	return &GetMeldescheinResponseBody{}
}

type GetMeldescheinResponseBody struct {
	XMLName xml.Name `xml:"holeMeldeschein"`

	Identifikation  Identifikation  `xml:"identifikation"`
	Fehlermeldungen Fehlermeldungen `xml:"fehlermeldungen"`
	Meldeschein     Meldeschein     `xml:"meldeschein"`
}

func (r *GetMeldeschein) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetMeldeschein) Do() (GetMeldescheinResponseBody, error) {
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
