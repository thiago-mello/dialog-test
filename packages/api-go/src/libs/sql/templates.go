package sql

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/leandro-andrade-candido/api-go/src/config"
)

var sqlTemplates *template.Template
var shouldGetTemplates = true

func init() {
	templatesPath := config.GetString("sql.templates.path")

	var err error
	sqlTemplates, err = template.ParseGlob(templatesPath)
	if err != nil {
		shouldGetTemplates = false
	}
}

// GetSql executes a named SQL template with provided data and returns the resulting SQL string
//
// Parameters:
//   - name: The name of the SQL template to lookup and execute
//   - data: Data to be used when executing the template
//
// Returns:
//   - string: The executed SQL template as a string
//   - error: Error if templates couldn't be loaded, template not found, or execution fails
func GetSql(name string, data any) (string, error) {
	if !shouldGetTemplates {
		return "", errors.New("could not load templates")
	}

	template := sqlTemplates.Lookup(name)
	if template == nil {
		return "", errors.New("template not found")
	}

	var buffer bytes.Buffer
	err := template.Execute(&buffer, data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
