package util

import (
	"github.com/buger/jsonparser"
	"github.com/edgedagency/mulungu/logger"
	"golang.org/x/net/context"
)

//JSONGetInt returns int64 value from json data
func JSONGetInt(ctx context.Context, subject map[string]interface{}, keys ...string) int64 {
	data := MapInterfaceToJSONBytes(subject)

	value, err := jsonparser.GetInt(data, keys...)
	if err != nil {
		logger.Errorf(ctx, "JSON util", "failed to retrieve key, %#v error:%s", keys, err.Error())
		return 0
	}

	return value
}

//JSONGetFloat returns float64 value from json data
func JSONGetFloat(ctx context.Context, subject map[string]interface{}, keys ...string) float64 {
	data := MapInterfaceToJSONBytes(subject)

	value, err := jsonparser.GetFloat(data, keys...)
	if err != nil {
		logger.Errorf(ctx, "JSON util", "failed to retrieve key, %#v error:%s", keys, err.Error())
		return 0.0
	}

	return value
}

//JSONGetString returns string value from json data
func JSONGetString(ctx context.Context, subject map[string]interface{}, keys ...string) string {
	data := MapInterfaceToJSONBytes(subject)

	value, err := jsonparser.GetString(data, keys...)
	if err != nil {
		logger.Errorf(ctx, "JSON util", "failed to retrieve key, %#v error:%s", keys, err.Error())
		return ""
	}

	return value
}

//JSONGetBoolean returns boolean value from json data
func JSONGetBoolean(ctx context.Context, subject map[string]interface{}, keys ...string) bool {
	data := MapInterfaceToJSONBytes(subject)

	value, err := jsonparser.GetBoolean(data, keys...)
	if err != nil {
		logger.Errorf(ctx, "JSON util", "failed to retrieve key, %#v error:%s", keys, err.Error())
		return false
	}

	return value
}
