package mulungu

import (
	"bytes"
	"html/template"

	"github.com/actdid/mulungu/util"
)

//TemplateParseHTMLFile reads HTML file and parses
func TemplateParseHTMLFile(name, path string, data interface{}) (processedTemplate string, err error) {

	fileContents, fileErr := util.FileRead(path)
	if fileErr != nil {
		return "", fileErr
	}

	buffer := bytes.NewBufferString(processedTemplate)
	tmpl := template.Must(template.New(name).Parse(string(fileContents)))
	err = tmpl.Execute(buffer, data)

	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
