package mulungu

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/clbanning/mxj"
)

//Response response object
type Response map[string]interface{}

//NewResponse create a new response
func NewResponse() Response {
	return make(Response)
}

//Add add value to response map object
func (r Response) Add(key string, value interface{}) Response {
	r[key] = value

	return r
}

//Get add value to response map object
func (r Response) Get(key string) interface{} {
	return r[key]
}

//JSON converts to json presentation
func (r Response) JSON() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return data
}

//XML converts to XML presentation
func (r Response) XML() []byte {
	mv := mxj.Map(r)
	data, xmlErr := mv.Xml()
	if xmlErr != nil {
		panic(xmlErr)
	}

	return data
}

//Format formats response based on content type
//
// For XML output header is attached.
func (r Response) Format(contentType string) []byte {
	switch strings.ToLower(strings.Replace(contentType, " ", "", -1)) {
	case "application/xml", "application/xml;charset=utf-8", "text/xml;charset=utf-8", "text/xml":
		return []byte(xml.Header + string(r.XML()))
	case "application/json", "application/json;charset=utf-8":
		return r.JSON()
	}
	//if all else fails dump string
	return []byte(fmt.Sprintf("Unable to format content for request, content type %s", contentType))
}
