package util

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/actdid/mulungu/constant"
	"github.com/actdid/mulungu/logger"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

//RuleExecute submits http request to rule for execution
func RuleExecute(ctx context.Context, region, namespace, scope string, data map[string]interface{}) (map[string]interface{}, error) {
	client := urlfetch.Client(ctx)

	fact := make(map[string]interface{})

	fact["scope"] = scope
	fact["namespace"] = namespace
	fact["fact"] = data
	// factBytes, factBytesErr := InterfaceToByte(fact)
	// if factBytesErr != nil {
	// 	logger.Errorf(ctx, "rule util", "fact to byte error %s", factBytesErr.Error())
	// 	return nil, factBytesErr
	// }
	//FIXME: consider turn below code into a utility
	request, requestError := http.NewRequest(http.MethodPost,
		CloudFunctionGetPath(region, appengine.AppID(ctx), "rule"), bytes.NewReader([]byte(MapInterfaceToJSONString(fact))))

	request.Header.Set(constant.HeaderContentType, "application/json")

	if requestError != nil {
		logger.Errorf(ctx, "rule util", "request init error %s", requestError.Error())
		return nil, requestError
	}

	dumpedRequest, _ := httputil.DumpRequest(request, true)
	logger.Debugf(ctx, "rule util", "Request %s", string(dumpedRequest))

	clientDoResponse, clientDoError := client.Do(request)

	dumpedResponse, _ := httputil.DumpResponse(clientDoResponse, true)
	logger.Debugf(ctx, "rule util", "Response %s", string(dumpedResponse))

	if clientDoError != nil {
		logger.Errorf(ctx, "rule util", "request execution error %s", clientDoError.Error())
		return nil, clientDoError
	}

	reponseData, responseError := ResponseToMap(clientDoResponse)
	if responseError != nil {
		logger.Errorf(ctx, "rule util", "request data mapping error %s", responseError.Error())
		return nil, responseError
	}

	return reponseData, nil
}

//CloudFunctionGetPath returns path to a cloud function
func CloudFunctionGetPath(region, projectID, function string) string {
	return fmt.Sprintf("%s://%s-%s.cloudfunctions.net/%s", "https", region, projectID, function)
}
