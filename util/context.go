package util

import (
	"net/http"

	"github.com/actdid/mulungu-go/constant"
	"github.com/actdid/mulungu-go/logger"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ContextAppEngine updates appengine context with namespace from header
func ContextAppEngine(r *http.Request) context.Context {

	context := appengine.NewContext(r)
	//wrap context in namespace if namespace is on request
	namespace := r.Header.Get(constant.HeaderNamespace)
	if namespace != "" {
		logger.Debugf(context, "appengine server", "wrapping context with namespace %s", namespace)
		wrappedContext, wrappingNamespaceError := appengine.Namespace(context, namespace)
		if wrappingNamespaceError != nil {
			logger.Criticalf(context, "appengine server", "failed to wrap namespace in context %s", wrappingNamespaceError.Error())
		} else {
			return wrappedContext
		}
	}

	return context

}
