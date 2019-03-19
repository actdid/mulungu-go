package provider

import (
	"net/http"
	"strings"

	"github.com/edgedagency/mulungu/constant"
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
)

//ArangodbDataProvider this is a dataprovider which helps with interacting with arangodb
type ArangodbDataProvider struct {
	Context   context.Context
	Namespace string
}

//NewArangodbDataProvider returns an ArangodbDataProvider
func NewArangodbDataProvider(ctx context.Context, namespace string) *ArangodbDataProvider {
	return &ArangodbDataProvider{Context: ctx, Namespace: namespace}
}

//Save save record
func (dp *ArangodbDataProvider) Save(collectionName string, data []byte) (map[string]interface{}, error) {

	response, responseError := dp.execute(collectionName, http.MethodPost, data, nil, nil)
	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return nil, responseError
	}

	responseMap, responseMapError := util.ResponseToMap(response)
	if responseMapError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "record conversation to map failed %s", responseMapError.Error())
		return nil, responseMapError
	}

	logger.Debugf(dp.Context, "Arangodb Data Provider", "record saved %#v", responseMap)

	return responseMap, nil
}

//Update update record
func (dp *ArangodbDataProvider) Update(collectionName string, id string, data []byte) (map[string]interface{}, error) {

	response, responseError := dp.execute(collectionName, http.MethodPut, data, nil, []string{id})
	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return nil, responseError
	}

	responseMap, responseMapError := util.ResponseToMap(response)
	if responseMapError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "record conversation to map failed %s", responseMapError.Error())
		return nil, responseMapError
	}

	logger.Debugf(dp.Context, "Arangodb Data Provider", "record updated %#v", responseMap)

	return responseMap, nil
}

//Delete delete record
func (dp *ArangodbDataProvider) Delete(collectionName string, id string) (map[string]interface{}, error) {

	response, responseError := dp.execute(collectionName, http.MethodDelete, nil, nil, []string{id})
	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return nil, responseError
	}

	responseMap, responseMapError := util.ResponseToMap(response)
	if responseMapError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "record conversation to map failed %s", responseMapError.Error())
		return nil, responseMapError
	}

	logger.Debugf(dp.Context, "Arangodb Data Provider", "record deleted %#v", responseMap)

	return responseMap, nil
}

//Find find record based on provided identifier
func (dp *ArangodbDataProvider) Find(collectionName string, id string) (map[string]interface{}, error) {
	response, responseError := dp.execute(collectionName, http.MethodGet, nil, nil, []string{id})
	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return nil, responseError
	}

	responseMap, responseMapError := util.ResponseToMap(response)
	if responseMapError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "record conversation to map failed %s", responseMapError.Error())
		return nil, responseMapError
	}

	logger.Debugf(dp.Context, "Arangodb Data Provider", "record found %#v", responseMap)

	return responseMap, nil
}

//FindAll search for and obtain records
func (dp *ArangodbDataProvider) FindAll(collectionName string, searchParams map[string]string) ([]interface{}, error) {

	response, responseError := dp.execute(collectionName, http.MethodGet, nil, searchParams, nil)

	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return nil, responseError
	}

	reponseMap, responseMapErr := util.ToInterfaceSlice(response.Body)
	if responseMapErr != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "failed to process reponse body %s", responseMapErr.Error())
		return nil, responseMapErr
	}

	logger.Debugf(dp.Context, "Arangodb Data Provider", "record found %#v", reponseMap)

	return reponseMap, nil
}

//Count return total number of records that match filter
func (dp *ArangodbDataProvider) Count(collectionName string, filter string) (int, error) {

	response, responseError := dp.execute(collectionName, http.MethodGet, nil, map[string]string{"filter": filter}, []string{"count"})

	if responseError != nil {
		logger.Errorf(dp.Context, "Arangodb Data Provider", "execution error %s", responseError.Error())
		return 0, responseError
	}

	return util.HTTPBodyAsInt64(response.Body), nil

}

func (dp *ArangodbDataProvider) addQueryParam(queryParams map[string]string, key string, value interface{}) map[string]string {
	return nil
}

func (dp *ArangodbDataProvider) execute(collectionName, method string, data []byte, searchParams map[string]string, pathParams []string) (*http.Response, error) {

	logger.Debugf(dp.Context, "Arangodb Data Provider", "executing request searchParams: %#v pathParams: %#v", searchParams, pathParams)

	//FIXME: util needed to construct a path
	return util.HTTPExecute(dp.Context, method, util.HTTPRequestURL(dp.Context,
		strings.Join([]string{"https://datastore-dot-ibudo-console.appspot.com/api/v1/arangodb", "datastore", collectionName}, "/"),
		pathParams, searchParams), map[string]string{constant.HeaderNamespace: dp.Namespace,
		constant.HeaderContentType: "application/json; charset=utf-8"},
		data,
		searchParams)
}
