package enova

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-enova/" + libraryVersion
	mediaType      = "application/soap+xml;charset=UTF-8"
	charset        = "utf-8"
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	dbName             string
	dbUsername         string
	dbPassword         string
	externalSystemGUID string
	tableName          string
	schemaName         string

	// User agent for client
	userAgent string

	mediaType string
	charset   string

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) DBName() string {
	return c.dbName
}

func (c *Client) SetDBName(dbName string) {
	c.dbName = dbName
}

func (c Client) DBUsername() string {
	return c.dbUsername
}

func (c *Client) SetDBUsername(dbUsername string) {
	c.dbUsername = dbUsername
}

func (c Client) DBPassword() string {
	return c.dbPassword
}

func (c *Client) SetDBPassword(dbPassword string) {
	c.dbPassword = dbPassword
}

func (c Client) ExternalSystemGUID() string {
	return c.externalSystemGUID
}

func (c *Client) SetExternalSystemGUID(externalSystemGUID string) {
	c.externalSystemGUID = externalSystemGUID
}

func (c Client) TableName() string {
	return c.tableName
}

func (c *Client) SetTableName(tableName string) {
	c.tableName = tableName
}

func (c Client) SchemaName() string {
	return c.schemaName
}

func (c *Client) SetSchemaName(schemaName string) {
	c.schemaName = schemaName
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to xml
	buf := new(bytes.Buffer)
	buf.Write([]byte(xml.Header))

	if req.RequestBodyInterface() != nil {
		soapRequest := NewRequestEnvelope()
		soapRequest.Body.ActionBody = req.RequestBodyInterface()

		enc := xml.NewEncoder(buf)
		enc.Indent("", "  ")
		err := enc.Encode(soapRequest)
		if err != nil {
			return nil, err
		}

		err = enc.Flush()
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	// r.Header.Add("SOAPAction", fmt.Sprintf("http://tempuri.org/RLXSOAP19/RLXSOAP19/%s", req.SOAPAction()))
	// r.Header.Add("SOAPAction", "urn:oracle:getKonfigurationsListe")
	// r.Header.Add("SOAPAction", "urn:oracle:holeMeldescheine")

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is xml decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	soapResponse := ResponseEnvelope{
		// Header: Header{},
		Body: Body{
			ActionBody: body,
		},
	}

	soapError := SoapError{}

	err = c.Unmarshal(httpResp.Body, &soapResponse, &soapError)
	if err != nil {
		return httpResp, err
	}

	if soapError.Body.Fault.FaultCode != "" || soapError.Body.Fault.FaultString != "" {
		return httpResp, &ErrorResponse{Response: httpResp, Err: soapError}
	}

	// if len(errorResponse.Messages) > 0 {
	// 	return httpResp, errorResponse
	// }

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := xml.NewDecoder(r)
		err := dec.Decode(v)
		if err != nil {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a xml response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	if r.StatusCode == 401 {
		return errors.New("401: Unauthorized")
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert xml to struct
	err = xml.Unmarshal(data, &errorResponse)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// <?xml version="1.0" encoding="UTF-8"?>
// <S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
//   <S:Body>
//     <ns0:holeMeldescheineResponse xmlns:ns0="http://www.avs.meldeschein.de/ns/">
//       <fehlermeldungen>
//         <fehler>
//           <code>10000</code>
//           <beschreibung>Unerwarteter technischer Fehler!</beschreibung>
//           <bezug>Meldeschein-Buchungsnummer: 10820</bezug>
//         </fehler>
//       </fehlermeldungen>
//     </ns0:holeMeldescheineResponse>
//   </S:Body>
// </S:Envelope>

// <?xml version="1.0" encoding="UTF-8"?>
// <S:Envelope xmlns:S="http://schemas.xmlsoap.org/soap/envelope/">
//   <S:Body>
//     <ns0:Fault xmlns:ns0="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://www.w3.org/2003/05/soap-envelope">
//       <faultcode>ns0:Client</faultcode>
//       <faultstring>Verteilungsmethode für {}holeMeldescheineRequest kann nicht gefunden werden</faultstring>
//     </ns0:Fault>
//   </S:Body>
// </S:Envelope>

type SoapError struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Fault struct {
			FaultCode   string `xml:"faultcode"`
			FaultString string `xml:"faultstring"`
		} `xml:"Fault"`
	} `xml:"Body"`
}

func (e SoapError) Error() string {
	return fmt.Sprintf("%s: %s", e.Body.Fault.FaultCode, e.Body.Fault.FaultString)
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// HTTP status code
	Err error
}

func (r *ErrorResponse) Error() string {
	return r.Err.Error()
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
