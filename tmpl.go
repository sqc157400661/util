package util

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"text/template"
)

// TrimmedLines Given a multi-line string, this function removes leading
// spaces from every line.
// It also removes the first line, if it is empty
func TrimmedLines(s string) string {
	// matches the start of the text followed by an EOL
	re := regexp.MustCompile(`(?m)\A\s*$`)
	s = re.ReplaceAllString(s, "")

	re = regexp.MustCompile(`(?m)^\t\t`)
	s = re.ReplaceAllString(s, "")
	return s
}

// Returns true if a StringMap (or an inner StringMap) contains a given key
func hasKey(sm map[string]interface{}, wantedKey string) bool {
	for key, value := range sm {
		if key == wantedKey {
			return true
		}
		valType := reflect.TypeOf(value)
		if valType == reflect.TypeOf(map[string]interface{}{}) {
			innerSm := value.(map[string]interface{})
			return hasKey(innerSm, wantedKey)
		}
		if valType == reflect.TypeOf([]map[string]interface{}{}) {
			innerSm := value.([]map[string]interface{})
			for _, ism := range innerSm {
				if hasKey(ism, wantedKey) {
					return true
				}
			}
		}
	}
	return false
}

// GetVarsFromTemplate Gets a list of all variables mentioned in a template
func GetVarsFromTemplate(tmpl string) []string {
	var varList []string

	reTemplateVar := regexp.MustCompile(`\{\{\.([^{]+)\}\}`)
	captureList := reTemplateVar.FindAllStringSubmatch(tmpl, -1)
	if len(captureList) > 0 {
		for _, capture := range captureList {
			varList = append(varList, capture[1])
		}
	}
	return varList
}

// SafeTemplateFill passed template string is formatted using its operands and returns the resulting string.
// It checks that the data was safely initialized
func SafeTemplateFill(tmpl string, data map[string]interface{}) (string, error) {
	if len(data) == 0 {
		return "", errors.New("no data")
	}
	// First, we get all variables in the pattern {{.VarName}}
	varList := GetVarsFromTemplate(tmpl)
	if len(varList) > 0 {
		for _, capture := range varList {
			// For each variable in the template text, we look whether it is
			// in the map
			if !hasKey(data, capture) {
				//fmt.Printf("### >>> %#v<<<\n", data)
				return "", fmt.Errorf("data field '%s'  was not initialized ",
					capture)
			}
		}
	}
	/**/
	// Creates a template
	processTemplate := template.Must(template.New("tmp").Parse(tmpl))
	buf := &bytes.Buffer{}

	// If an error occurs, returns an empty string
	if err := processTemplate.Execute(buf, data); err != nil {
		return "", err
	}
	// Returns the populated template
	return buf.String(), nil
}
