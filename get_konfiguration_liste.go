package meldeschein

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-avs-meldeschein/utils"
)

func (c *Client) NewGetConfigurationList() GetConfigurationList {
	r := GetConfigurationList{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetConfigurationList struct {
	client      *Client
	queryParams *GetConfigurationListQueryParams
	pathParams  *GetConfigurationListPathParams
	method      string
	headers     http.Header
	requestBody GetConfigurationListBody
}

func (r GetConfigurationList) NewQueryParams() *GetConfigurationListQueryParams {
	return &GetConfigurationListQueryParams{}
}

type GetConfigurationListQueryParams struct {
}

func (p GetConfigurationListQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetConfigurationList) QueryParams() QueryParams {
	return r.queryParams
}

func (r GetConfigurationList) NewPathParams() *GetConfigurationListPathParams {
	return &GetConfigurationListPathParams{}
}

type GetConfigurationListPathParams struct {
}

func (p *GetConfigurationListPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetConfigurationList) PathParams() PathParams {
	return r.pathParams
}

func (r *GetConfigurationList) SetMethod(method string) {
	r.method = method
}

func (r *GetConfigurationList) Method() string {
	return r.method
}

func (r GetConfigurationList) NewRequestBody() GetConfigurationListBody {
	return GetConfigurationListBody{}
}

type GetConfigurationListBody struct {
	XMLName        xml.Name       `xml:"ns:configuration-lists"`
	Identifikation Identifikation `xml:"identifikation"`
}

func (r *GetConfigurationList) RequestBody() *GetConfigurationListBody {
	return &r.requestBody
}

func (r *GetConfigurationList) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetConfigurationList) SetRequestBody(body GetConfigurationListBody) {
	r.requestBody = body
}

func (r *GetConfigurationList) NewResponseBody() *GetConfigurationListResponseBody {
	return &GetConfigurationListResponseBody{}
}

type GetConfigurationListResponseBody struct {
	XMLName xml.Name `xml:"configuration-lists"`

	Identifikation struct {
		Erzeugung     string `xml:"erzeugung"`
		Schnittstelle string `xml:"schnittstelle"`
		Kurverwaltung string `xml:"kurverwaltung"`
		Benutzerid    string `xml:"benutzerid"`
		Verarbeitung  string `xml:"verarbeitung"`
		Version       string `xml:"version"`
	} `xml:"identifikation"`
	Fehlermeldungen struct {
		Fehler struct {
			Code         string `xml:"code"`
			Beschreibung string `xml:"beschreibung"`
		} `xml:"fehler"`
	} `xml:"fehlermeldungen"`
	Konfigliste struct {
		Konfigdatensatz []struct {
			ID    string `xml:"id"`
			Text1 string `xml:"text1"`
			Text2 string `xml:"text2"`
			Text3 string `xml:"text3"`
		} `xml:"konfigdatensatz"`
	} `xml:"konfigliste"`
}

func (r *GetConfigurationList) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *GetConfigurationList) Do() (GetConfigurationListResponseBody, error) {
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
