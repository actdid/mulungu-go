package api

import "golang.org/x/net/context"

var configuration *Configuration

func init() {
	configuration = &Configuration{}
}

//Configuration configuration management struct
type Configuration struct {
}

//ConfigurationGet return based on pointer variable
func ConfigurationGet(ctx context.Context, namespace string, key string, defaultValue interface{}) interface{} {
	return defaultValue
}

//Get return configuration value based in key
func (c *Configuration) Get(ctx context.Context, namespace, key, defaultValue interface{}) interface{} {
	return defaultValue
}

//Set set configuration value
func (c *Configuration) Set(ctx context.Context, namespace, key, value interface{}) interface{} {
	return value
}

//ConfigurationGetString return string value
func ConfigurationGetString(ctx context.Context, namespace, key, defaultValue interface{}) string {
	return configuration.GetString(ctx, namespace, key, defaultValue)
}

//GetString return string value
func (c *Configuration) GetString(ctx context.Context, namespace, key, defaultValue interface{}) string {
	return c.Get(ctx, namespace, key, defaultValue).(string)
}
