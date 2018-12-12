package transformer

import (
	"regexp"
	"strings"
)

var (
	templateItemRegex *regexp.Regexp
)

func Project(item DataObject, projector Projector) DataObject {
	if templateItemRegex == nil {
		templateItemRegex = regexp.MustCompile(`\${{[\w\d_]+}}`)
	}
	result := DataObject{}
	for _, projection := range projector {
		result[projection.OutKey] = populateTemplate(projection.Template, item)
	}
	return result
}

func populateTemplate(template string, object DataObject) string {
	templateItems := templateItemRegex.FindAllString(template, -1)
	for _, templateItem := range templateItems {
		key := getTemplateItemKey(templateItem)
		value := object[key]
		template = strings.Replace(template, templateItem, value, -1)
	}
	return template
}

func getTemplateItemKey(templateItem string) string {
	templateItem = strings.TrimPrefix(templateItem, "${{")
	templateItem = strings.TrimSuffix(templateItem, "}}")
	return templateItem

}
