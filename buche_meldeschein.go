package meldeschein

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-avs-meldeschein/utils"
)

func (c *Client) NewPostMeldeschein() PostMeldeschein {
	r := PostMeldeschein{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PostMeldeschein struct {
	client      *Client
	queryParams *PostMeldescheinQueryParams
	pathParams  *PostMeldescheinPathParams
	method      string
	headers     http.Header
	requestBody PostMeldescheinRequestBody
}

func (r PostMeldeschein) NewQueryParams() *PostMeldescheinQueryParams {
	return &PostMeldescheinQueryParams{}
}

type PostMeldescheinQueryParams struct {
}

func (p PostMeldescheinQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostMeldeschein) QueryParams() QueryParams {
	return r.queryParams
}

func (r PostMeldeschein) NewPathParams() *PostMeldescheinPathParams {
	return &PostMeldescheinPathParams{}
}

type PostMeldescheinPathParams struct {
}

func (p *PostMeldescheinPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostMeldeschein) PathParams() PathParams {
	return r.pathParams
}

func (r *PostMeldeschein) SetMethod(method string) {
	r.method = method
}

func (r *PostMeldeschein) Method() string {
	return r.method
}

func (r PostMeldeschein) NewRequestBody() PostMeldescheinRequestBody {
	return PostMeldescheinRequestBody{
		Identifikation: Identifikation{
			Erzeugung:     Date{time.Now()},
			Kurverwaltung: 0,
			BenutzerID:    1,
			Verarbeitung:  "BUCHEN",
		},
	}
}

type PostMeldescheinRequestBody struct {
	XMLName        xml.Name       `xml:"ns:meldescheine"`
	Identifikation Identifikation `xml:"identifikation"`
	Meldeschein    Meldeschein    `xml:"ns:meldeschein"`
}

func (r *PostMeldeschein) RequestBody() *PostMeldescheinRequestBody {
	return &r.requestBody
}

func (r *PostMeldeschein) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostMeldeschein) SetRequestBody(body PostMeldescheinRequestBody) {
	r.requestBody = body
}

func (r *PostMeldeschein) NewResponseBody() *PostMeldescheinResponseBody {
	return &PostMeldescheinResponseBody{}
}

type PostMeldescheinResponseBody struct {
	XMLName xml.Name `xml:"meldescheine"`

	Identifikation  Identifikation  `xml:"identifikation"`
	Fehlermeldungen Fehlermeldungen `xml:"fehlermeldungen>fehler"`
}

func (r *PostMeldeschein) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *PostMeldeschein) Do() (PostMeldescheinResponseBody, error) {
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
		return *responseBody, err
	}

	if len(responseBody.Fehlermeldungen) > 0 {
		return *responseBody, responseBody.Fehlermeldungen[0]
	}

	return *responseBody, nil
}
