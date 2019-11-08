package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/actdid/mulungu/constant"
	"github.com/actdid/mulungu/logger"
	"github.com/actdid/mulungu/util"
	"golang.org/x/net/context"
)

const defaultNameSpace = "ibudo.xyz"

func pubSubData(ctx context.Context, r *http.Request) (map[string]interface{}, error) {
	data, err := util.ToMapStringInterface(r.Body)
	logger.Debugf(ctx, "middleware namespace", "topic data %#v", data)
	mappedData := make(map[string]interface{})
	if err == nil {

		mappedData["subscription"] = util.PubSubTopicSubscription(ctx, data)
		mappedData["attributes"] = util.MapInterfaceToMapString(util.PubSubTopicAttributes(ctx, data))
		mappedData["data"] = util.PubSubTopicData(ctx, data)

		logger.Debugf(ctx, "middleware namespace", "mapped topic data %#v", mappedData)

		return mappedData, nil
	}

	return nil, err
}

//Namespace this handler determins and set appropriate namespace on re
func Namespace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := util.ContextAppEngine(r)

		logger.Debugf(ctx, "middleware namespace", "determining namespace for request, current namespace: %s ", r.Header.Get(constant.HeaderNamespace))

		if strings.ToLower(util.HTTPGetQueryValue(r, "source", "")) == "pubsub" {
			logger.Debugf(ctx, "middleware namespace", "topic subscription, obtaining namespace from subscription")
			//FIXME: quick utility function to return request data, need to extract namespace and other information
			data, err := pubSubData(ctx, r)
			if err == nil {
				r.Header.Set(constant.HeaderPubSub, "true")
				subscriptionNamesapcePart := strings.ToLower(data["subscription"].([]string)[0])
				logger.Debugf(ctx, "middleware namespace", "topic subscription namespace %s", subscriptionNamesapcePart)
				r.Header.Set(constant.HeaderNamespace, subscriptionNamesapcePart)

				//golang looses body once read, it must be rewritten
				r.Body = ioutil.NopCloser(bytes.NewReader(util.MapInterfaceToJSONBytes(data)))
				r.ContentLength = int64(len(data))
			}
			if err != nil {
				logger.Errorf(ctx, "middleware namespace", "failed to decode request, to obtain namespace %s", err.Error())
			}
		}

		if r.Header.Get(constant.HeaderNamespace) != "" {
			r.Header.Set(constant.HeaderNamespace, util.SetEnvironmentOnNamespace(ctx, r.Header.Get(constant.HeaderNamespace), r))
			logger.Debugf(ctx, "middleware namespace", "Request has namespace %s", r.Header.Get(constant.HeaderNamespace))
		} else if util.SetNamespace(ctx, r) != "" {
			logger.Debugf(ctx, "middleware namespace", "obtained namespace from url, adding environment")
			r.Header.Set(constant.HeaderNamespace, util.SetEnvironmentOnNamespace(ctx, util.SetNamespace(ctx, r), r))
		}

		logger.Debugf(ctx, "middleware namespace", "request namespace set to %s ", r.Header.Get(constant.HeaderNamespace))

		next.ServeHTTP(w, r)
	})
}
