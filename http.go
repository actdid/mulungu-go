package mulungu

import (
	"bytes"
	"strings"
	"time"

	"golang.org/x/net/context"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/actdid/mulungu/logger"
	"github.com/actdid/mulungu/util"
	"google.golang.org/appengine/urlfetch"
)

//HTTPResponse is a structure which is used to store http request responses
type HTTPResponse struct {
	context  context.Context
	Response *http.Response
	Error    error
}

//HTTPRequest is a structure which will be used to represent a connection
type HTTPRequest struct {
	context    context.Context
	host       string
	username   string
	password   string
	httpClient *http.Client
	headers    map[string]string
	response   *HTTPResponse
	secured    bool
	schema     string
}

//GenerateGoogleServiceHost this returns an app spot host
func GenerateGoogleServiceHost(host, service string) string {
	googleServiceHost := fmt.Sprintf("%s-dot-%s", service, host)
	return googleServiceHost
}

// NewHTTPRequest creates connection and returns &pointer to refrence to this connection.
func NewHTTPRequest(ctx context.Context, schema, host, username, password string, secured bool, headers map[string]string) *HTTPRequest {

	logger.Debugf(ctx, "http request", "creating new request schema:%s host:%s headers:%v", schema, host, headers)
	ctxWithDeadLine, _ := context.WithDeadline(ctx, time.Now().Add(30*time.Second))

	return &HTTPRequest{context: ctx,
		host:       schema + "://" + host,
		username:   username,
		password:   password,
		headers:    headers,
		secured:    secured,
		httpClient: urlfetch.Client(ctxWithDeadLine)}
}

// SendJSON construct a json request body for this request and sends to configured http endpoint
func (httpRequest *HTTPRequest) SendJSON(httpMethod, requestPath string, requestBody interface{}) *HTTPResponse {
	marshalledRequestBody, err := json.Marshal(requestBody)
	requestURL := httpRequest.constructRequestURL(requestPath)

	logger.Debugf(httpRequest.context, "http request", "preparing request details, method %s, requestPath %s, requestURL %s", httpMethod, requestPath, requestURL)
	request, err := http.NewRequest(httpMethod, requestURL, bytes.NewBuffer(marshalledRequestBody))

	if err != nil {
		logger.Errorf(httpRequest.context, "http request", "failed to create http request %s", err.Error())
		return &HTTPResponse{context: httpRequest.context, Response: nil, Error: err}
	}

	httpRequest.setHeaders(request)
	httpRequest.setAuthentication(request)

	logger.Debugf(httpRequest.context, "http request", "sending request %#v to service", request)

	httpClientResponse, httpClientResponseErr := httpRequest.httpClient.Do(request)
	logger.Debugf(httpRequest.context, "http request", "processing response from service call")

	if httpClientResponseErr != nil {
		logger.Errorf(httpRequest.context, "http request", "request failed, error %s", httpClientResponseErr.Error())
	}

	return &HTTPResponse{context: httpRequest.context, Response: httpClientResponse, Error: httpClientResponseErr}
}

func (httpRequest *HTTPRequest) constructRequestURL(requestPath string) string {
	if strings.HasPrefix(requestPath, "/") == false {
		if strings.HasSuffix(httpRequest.host, "/") == false {
			requestPath = "/" + requestPath
		}
	}

	return fmt.Sprintf("%s%s", httpRequest.host, requestPath)
}

func (httpRequest *HTTPRequest) setHeaders(request *http.Request) {
	logger.Debugf(httpRequest.context, "http request", "setting headers")
	for key, value := range httpRequest.headers {
		logger.Debugf(httpRequest.context, "http request", "key %s value %s", key, value)
		request.Header.Set(key, value)
	}
}

func (httpRequest *HTTPRequest) setAuthentication(request *http.Request) {
	//FIXME: if connection is secured look for common security headers, properties before proceeding
	if httpRequest.secured {
		if httpRequest.username != "" && httpRequest.password != "" {
			request.SetBasicAuth(httpRequest.username, httpRequest.password)
		}
	}
}

//HasErrors determines in HTTPResponse has errors from executed request
func (httpResponse *HTTPResponse) HasErrors() bool {
	logger.Debugf(httpResponse.context, "http request", "checking is http request has errornous response")

	if httpResponse.Response == nil {
		logger.Criticalf(httpResponse.context, "http request", "no response from executed HttpRequest")
		return true
	}

	logger.Debugf(httpResponse.context, "http request", "response status %s", httpResponse.Response.Status)
	logger.Debugf(httpResponse.context, "http request", "response code %s", httpResponse.Response.StatusCode)

	if httpResponse.Error != nil {
		return true
	}

	logger.Debugf(httpResponse.context, "http request", "no errors found")
	return false
}

//JSON returns http request response to json map[string]interface{}, use this method if you are expecting json response
func (httpResponse *HTTPResponse) JSON() (response map[string]interface{}, err error) {
	logger.Debugf(httpResponse.context, "http request", "converting response to map/json")
	return util.JSONDecodeHTTPResponse(httpResponse.Response)
}
