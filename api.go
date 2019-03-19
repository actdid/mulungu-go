package mulungu

import "golang.org/x/net/context"

//API represenetation of an api
type API struct {
	context   context.Context
	namespace string
}

//Init initiates this api
func (s *API) Init(context context.Context, namespace string) {
	s.context = context
	s.namespace = namespace
}

//Namespace returns namespace
func (s *API) Namespace() string {
	return s.namespace
}

//Context returns context
func (s *API) Context() context.Context {
	return s.context
}
