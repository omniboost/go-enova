package cardxperts

import (
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/omniboost/go-cardxperts/utils"
)

func (c *Client) NewListGastarten() ListGastarten {
	r := ListGastarten{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ListGastarten struct {
	client      *Client
	queryParams *ListGastartenQueryParams
	pathParams  *ListGastartenPathParams
	method      string
	headers     http.Header
	requestBody ListGastartenBody
}

func (r ListGastarten) NewQueryParams() *ListGastartenQueryParams {
	return &ListGastartenQueryParams{}
}

type ListGastartenQueryParams struct {
}

func (p ListGastartenQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ListGastarten) QueryParams() QueryParams {
	return r.queryParams
}

func (r ListGastarten) NewPathParams() *ListGastartenPathParams {
	return &ListGastartenPathParams{}
}

type ListGastartenPathParams struct {
}

func (p *ListGastartenPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ListGastarten) PathParams() PathParams {
	return r.pathParams
}

func (r *ListGastarten) SetMethod(method string) {
	r.method = method
}

func (r *ListGastarten) Method() string {
	return r.method
}

func (r ListGastarten) NewRequestBody() ListGastartenBody {
	return ListGastartenBody{}
}

type ListGastartenBody struct {
	XMLName xml.Name `xml:"car:listlgastarten"`
}

func (r *ListGastarten) RequestBody() *ListGastartenBody {
	return &r.requestBody
}

func (r *ListGastarten) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *ListGastarten) SetRequestBody(body ListGastartenBody) {
	r.requestBody = body
}

func (r *ListGastarten) NewResponseBody() *ListGastartenResponseBody {
	return &ListGastartenResponseBody{}
}

type ListGastartenResponseBody struct {
	XMLName xml.Name `xml:"listlaendercodesResponse"`
}

func (r *ListGastarten) URL() *url.URL {
	u := r.client.GetEndpointURL("", r.PathParams())
	return &u
}

func (r *ListGastarten) Do() (ListGastartenResponseBody, error) {
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
