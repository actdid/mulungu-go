package util

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/actdid/mulungu-go/logger"
	"github.com/uniplaces/carbon"
)

//TemplateParse parses content data
func TemplateParse(ctx context.Context, source string, data interface{}) string {

	//custom functions
	funcMap := template.FuncMap{"date": func(value, format string) string {
		parsedDate, _ := time.Parse(carbon.RFC3339Format, value)
		logger.Debugf(ctx, "template parse util", "date: %s parsed: %s format: %s", value, parsedDate.String(), format)
		return parsedDate.Format(format)
	}, "trim": func(value string) string {
		return strings.TrimSpace(value)
	}}

	templateIdentifier := strings.Join([]string{"template", GenerateRandomCode(5, "")}, "")
	t, err := template.New(templateIdentifier).Funcs(funcMap).Option("missingkey=zero").Parse(source)

	logger.Debugf(ctx, "template parse util", "parsing template, templateIdentifier: %s data: %#v", templateIdentifier, data)

	if err != nil {
		logger.Errorf(ctx, "template parse util", "parse template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template %s error %s", source, err.Error())
	}

	var output bytes.Buffer

	if err := t.Execute(&output, data); err != nil {
		logger.Errorf(ctx, "template parse util", "execute template failed with error %s", err.Error())
		return fmt.Sprintf("failed to parse template, execution failed %s error %s", source, err.Error())
	}

	return output.String()
}
