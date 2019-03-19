package pubsub

import (
	"github.com/edgedagency/mulungu/logger"
	"github.com/edgedagency/mulungu/util"
	"golang.org/x/net/context"
)

//PushSubscriptionInfo type for working with pub/sub messages from google cloud
type PushSubscriptionInfo struct {
	data         map[string]interface{}
	Context      context.Context
	namespace    string
	subscription string
}

//NewPushSubscriptionInfo returns pointer to PushSubscription, use to interact with push subscription
func NewPushSubscriptionInfo(ctx context.Context, publishedData map[string]interface{}) *PushSubscriptionInfo {
	logger.Debugf(ctx, "pubsub", "building up Publish/Subscription Info, with published data %#v", publishedData)

	subscriptionInfo := SubscritpionInfo(ctx, publishedData)
	data := SubscritpionData(ctx, publishedData)

	instance := &PushSubscriptionInfo{Context: ctx, data: data, namespace: subscriptionInfo[0], subscription: subscriptionInfo[1]}

	return instance
}

//Subscription returns subscription part of message
func (ps *PushSubscriptionInfo) Subscription() string {
	return ps.subscription
}

//Namespace returns namespace part of message
func (ps *PushSubscriptionInfo) Namespace() string {
	return ps.namespace
}

//Data returns data publiched to topic
func (ps *PushSubscriptionInfo) Data() map[string]interface{} {
	return ps.data
}

//SubscritpionInfo returns namespace/subscription []string{namespace,subscription} from published data constructed in namespace matching
func SubscritpionInfo(ctx context.Context, data map[string]interface{}) []string {
	if subscription, ok := data["subscription"]; ok {
		logger.Debugf(ctx, "pubsub", "subscription %#v", data["subscription"])
		return util.InterfaceToStringSlice(subscription)
	}
	return nil
}

//SubscritpionData returns topic data as map[string]interface{} decoded JSON data
func SubscritpionData(ctx context.Context, data map[string]interface{}) map[string]interface{} {
	if subscritpionData, ok := data["data"]; ok {
		return subscritpionData.(map[string]interface{})
	}
	return nil
}
