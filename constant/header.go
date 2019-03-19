package constant

//Constant for http headers
const (
	HeaderProxyHost                 = "X-Proxy-Host"
	HeaderNamespace                 = "X-Namespace"
	HeaderKind                      = "X-Kind"
	HeaderRequestAuthorizationToken = "X-Request-Authorization-Token"
	HeaderAuthorization             = "Authorization"
	HeaderContentType               = "Content-Type"
	HeaderAuthorised                = "X-Authorised"
	HeaderAuthorisedRole            = "X-Authorised-Role"
	HeaderTenantID                  = "X-Authorised-Tenant-ID"
	HeaderEnvironment               = "X-Environment"
	HeaderServiceHost               = "X-Service-Host"
	HeaderRequestSource             = "X-Request-Source"
	HeaderPubSub                    = "X-PubSub"
	HeaderCron                      = "X-Appengine-Cron"
)
