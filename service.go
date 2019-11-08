package mulungu

import (
	"github.com/actdid/mulungu-go/core"
	"github.com/actdid/mulungu-go/logger"
	"github.com/actdid/mulungu-go/util"
	"golang.org/x/net/context"
)

//Service represenetation of a service
type Service struct {
	context   context.Context
	namespace string
	kind      string
}

//Init initiates this services
func (s *Service) Init(context context.Context, namespace, kind string) {
	s.context = context
	s.namespace = namespace
	s.kind = kind
}

//Kind returns kind
func (s *Service) Kind() string {
	return s.kind
}

//Namespace returns namespace
func (s *Service) Namespace() string {
	return s.namespace
}

//Context returns context
func (s *Service) Context() context.Context {
	return s.context
}

//Find returns service based on id
func (s *Service) Find(id string) (map[string]interface{}, error) {

	//2. save record
	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), nil)
	responseRecord, responseError := datastoreModel.Find(id)

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//FindAll finds all service from datastore
func (s *Service) FindAll(filter string) ([]interface{}, error) {

	logger.Debugf(s.Context(), "service service", "finding all %s", filter)

	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), nil)
	responseRecord, responseError := datastoreModel.FindAll(core.NewSearchParam().Add("filter", filter).AsMap())

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//Update returns service based on id
func (s *Service) Update(id string, record map[string]interface{}) (map[string]interface{}, error) {

	//2. save record
	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), record)
	responseRecord, responseError := datastoreModel.Update(id)

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//Delete returns service based on id
func (s *Service) Delete(id string) (map[string]interface{}, error) {

	//2. save record
	datastoreModel := core.NewDatastoreModel(s.Context(), s.Namespace(), s.Kind(), nil)
	responseRecord, responseError := datastoreModel.Delete(id)

	if responseError != nil {
		return nil, responseError
	}

	return responseRecord, nil
}

//Publish publish to topic
func (s *Service) Publish(topic string, record map[string]interface{}, attributes map[string]string) (string, error) {
	publishID, publishErr := util.PubSubTopicPublish(s.Context(), util.PubSubTopicID(s.Namespace(), topic), record, attributes)
	if publishErr != nil {
		logger.Errorf(s.Context(), "service service", "failed to publish record created %s", publishErr.Error())
		return "", publishErr
	}
	logger.Debugf(s.Context(), "service service", "published record created, publish id %s", publishID)
	return publishID, nil
}

//PublishNoNamespace meant to replace publish feature, records published already have a namespace, infer namespace
func (s *Service) PublishNoNamespace(topic string, record map[string]interface{}, attributes map[string]string) (string, error) {
	publishID, publishErr := util.PubSubTopicPublish(s.Context(), topic, record, attributes)
	if publishErr != nil {
		logger.Errorf(s.Context(), "service service", "failed to publish record created %s", publishErr.Error())
		return "", publishErr
	}
	logger.Debugf(s.Context(), "service service", "published record created, publish id %s", publishID)
	return publishID, nil
}
