package middleware

import (
	"net/http"

	"github.com/actdid/mulungu-go/constant"
	"github.com/actdid/mulungu-go/logger"

	"google.golang.org/appengine"
)

//Logging this handler determins and set appropriate namespace on re
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		logger.Debugf(ctx, "middleware logging", "logging request information")

		logger.Debugf(ctx, "middleware logging", "request headers")
		for header, value := range r.Header {
			logger.Debugf(ctx, "middleware logging", "Header: %s Value: %s", header, value)
		}

		logger.Debugf(ctx, "middleware logging", "request")
		logger.Debugf(ctx, "middleware logging", "BEFORE: URI:%s DefaultHostName:%s Host:%s Path:%s Referer:%s IP:%s Scheme:%s Method:%s X-Namespace:%s appId:%s datacenter:%s environment:%s proxy-host:%s authorised:%s roles:%s",
			r.URL.RequestURI(),
			appengine.DefaultVersionHostname(ctx),
			r.URL.Host,
			r.URL.Path,
			r.Referer(),
			r.RemoteAddr,
			r.URL.Scheme,
			r.Method,
			r.Header.Get(constant.HeaderNamespace),
			appengine.AppID(ctx),
			appengine.Datacenter(ctx),
			r.Header.Get(constant.HeaderEnvironment),
			r.Header.Get(constant.HeaderProxyHost),
			r.Header.Get(constant.HeaderAuthorised),
			r.Header.Get(constant.HeaderAuthorisedRole))

		next.ServeHTTP(w, r)

		logger.Debugf(ctx, "middleware logging", "AFTER: URI:%s DefaultHostName:%s Host:%s Path:%s Referer:%s IP:%s Scheme:%s Method:%s X-Namespace:%s appId:%s datacenter:%s environment:%s proxy-host:%s authorised:%s roles:%s",
			r.URL.RequestURI(),
			appengine.DefaultVersionHostname(ctx),
			r.URL.Host,
			r.URL.Path,
			r.Referer(),
			r.RemoteAddr,
			r.URL.Scheme,
			r.Method,
			r.Header.Get(constant.HeaderNamespace),
			appengine.AppID(ctx),
			appengine.Datacenter(ctx),
			r.Header.Get(constant.HeaderEnvironment),
			r.Header.Get(constant.HeaderProxyHost),
			r.Header.Get(constant.HeaderAuthorised),
			r.Header.Get(constant.HeaderAuthorisedRole))
	})
}
