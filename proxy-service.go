package mulungu

import (
	"net/http"

	"github.com/actdid/mulungu/logger"
	"github.com/actdid/mulungu/util"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ProxyService object used to make proxy services
type ProxyService struct {
	Context context.Context
	Schema  string
	Service string
	Path    string
}

//NewProxyService returns proxy service handler
func NewProxyService(ctx context.Context, schema, service, path string) *ProxyService {
	proxyService := &ProxyService{Context: ctx, Schema: schema, Service: service, Path: path}
	return proxyService
}

//SendHost sends request to provided host
func (ps *ProxyService) SendHost(host string, method string, data map[string]interface{}, username, password string, secured bool, headers map[string]string) (*HTTPResponse, error) {
	httpResponse := NewHTTPRequest(ps.Context, ps.Schema, GenerateGoogleServiceHost(host, ps.Service), username, password, secured, headers).SendJSON(method, ps.Path, data)

	if httpResponse.HasErrors() {
		logger.Errorf(ps.Context, "proxy service", "failed to processes/proxy request, error: %s", httpResponse.Error.Error())
		return httpResponse, httpResponse.Error
	}

	return httpResponse, nil
}

//Send used to send proxy request
func (ps *ProxyService) Send(method string, data map[string]interface{}, username, password string, secured bool, headers map[string]string) (*HTTPResponse, error) {
	serviceHost := GenerateGoogleServiceHost(appengine.DefaultVersionHostname(ps.Context), ps.Service)
	httpResponse := NewHTTPRequest(ps.Context, ps.Schema, serviceHost, username, password, secured, headers).SendJSON(method, ps.Path, data)

	if httpResponse.HasErrors() {
		logger.Errorf(ps.Context, "proxy service", "failed to processes/proxy request, error: %s", httpResponse.Error.Error())
		return httpResponse, httpResponse.Error
	}

	return httpResponse, nil
}

//ProxyHost shortcut method that executes a proxy request with a different host
func (ps *ProxyService) ProxyHost(host string, w http.ResponseWriter, r *http.Request) (*HTTPResponse, error) {
	data, dataErr := util.JSONDecodeHTTPRequest(r)
	if dataErr != nil {
		logger.Errorf(ps.Context, "proxy service", "failed to decode body, error : %s", dataErr.Error())
	}
	logger.Errorf(ps.Context, "proxy service", "firing off request, data:%v service:%s path:%s", data, ps.Service, ps.Path)
	return ps.SendHost(host, r.Method, data, "", "", false, map[string]string{"Content-Type": "application/json; charset=utf-8", "X-Namespace": r.Header.Get("Host")})
}

//Proxy shortcut method that executes a proxy request
func (ps *ProxyService) Proxy(w http.ResponseWriter, r *http.Request) (*HTTPResponse, error) {
	data, dataErr := util.JSONDecodeHTTPRequest(r)
	if dataErr != nil {
		logger.Errorf(ps.Context, "proxy service", "failed to decode body, error : %s", dataErr.Error())
	}
	logger.Debugf(ps.Context, "proxy service", "firing off request, data:%v service:%s path:%s", data, ps.Service, ps.Path)
	return ps.Send(r.Method, data, "", "", false, map[string]string{"Content-Type": "application/json; charset=utf-8", "X-Namespace": r.Header.Get("Host")})
}
