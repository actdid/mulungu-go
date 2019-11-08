package mulungu

import (
	"net/http"
	"strings"

	"github.com/actdid/mulungu-go/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//AppEngineServer representation of Google App Engine Server
var AppEngineServer *AppEngine

//Runnning indicates that server is running
var Runnning = false

func init() {
	AppEngineServer = NewAppEngine()
}

//AppEngine Google appengine server configurations
type AppEngine struct {
	router *mux.Router
	chain  alice.Chain
}

//NewAppEngine create a new appengine server
func NewAppEngine() *AppEngine {
	return &AppEngine{router: mux.NewRouter(), chain: alice.New(middleware.Logging, middleware.Namespace)}
}

//Start sets up http handler with register handlers
func (s *AppEngine) Start() {
	if Runnning == false {
		http.Handle("/", s.chain.Then(s.router))
		Runnning = true
	}
}

//Middleware registers middlewares
func (s *AppEngine) Middleware(middlwares ...alice.Constructor) {
	s.chain = s.chain.Append(middlwares...)
}

//Handler can be used to register a handler, handlers process information based on a path signature
func (s *AppEngine) Handler(path string, h http.Handler) *mux.Route {
	return s.router.Handle(path, h)
}

//HandlerFunc can be used to register a handler, handlers process information based on a path signature
func (s *AppEngine) HandlerFunc(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return s.router.HandleFunc(path, f)
}

// //Context returns context from request
// func (s *AppEngine) Context(r *http.Request) context.Context {
// 	context := appengine.NewContext(r)
// 	//wrap context in namespace if namespace is on request
// 	namespace := r.Header.Get(constant.HeaderNamespace)
// 	if namespace != "" {
// 		logger.Debugf(context, "appengine server", "wrapping context with namespace %s", namespace)
// 		wrappedContext, wrappingNamespaceError := appengine.Namespace(context, namespace)
// 		if wrappingNamespaceError != nil {
// 			logger.Criticalf(context, "appengine server", "failed to wrap namespace in context %s", wrappingNamespaceError.Error())
// 		} else {
// 			return wrappedContext
// 		}
// 	}
//
// 	return context
// }

//AppEngineServiceURL this returns an app spot host
func AppEngineServiceURL(host, service, version string) string {
	var hostURL = strings.Join([]string{service, host}, "-dot-")
	//e.g. https://v10032017t163649-dot-application-dot-ibudo-console.appspot.com/
	if version != "" {
		hostURL = strings.Join([]string{version, service, host}, "-dot-")
	}

	return hostURL
}
