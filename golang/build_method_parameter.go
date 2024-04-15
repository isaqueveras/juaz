package golang

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_method_parameter(juaz *grammar.Juaz, modelName, node string) string {
	var buf strings.Builder

	for _, value := range juaz.Entries {
		if value.Node == nil || value.Node.Name != node {
			continue
		}

		modelName := strcase.ToCamel(modelName)
		modelNameLower := strings.ToLower(string(modelName[0]))

		buf.WriteString("\n// NewParams ...")
		buf.WriteString(fmt.Sprintf("\nfunc (%s *%s) NewParams() {\n", modelNameLower, modelName))
		buf.WriteString(fmt.Sprintf("\t%s.parameters = &%s{}\n}\n", modelNameLower, strcase.ToCamel(node)))

		for _, entry := range value.Node.Entries {
			field := _build_type_field(entry.Field)
			method := strcase.ToCamel(entry.Field.Name)
			params := strcase.ToLowerCamel(entry.Field.Name)

			buf.WriteString("\n// WithParam" + method + " ...")
			buf.WriteString(fmt.Sprintf("\nfunc (%s *%s) WithParam%s(%s%s) {\n", modelNameLower, modelName, method, params, field))
			buf.WriteString(fmt.Sprintf("\t%s.parameters.%s = %s\n}\n", modelNameLower, method, params))
		}
	}

	return buf.String()
}
