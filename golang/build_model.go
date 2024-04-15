package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_model(juaz *grammar.Juaz) string {
	var buf strings.Builder

	for _, value := range juaz.Entries {
		if value.Node == nil {
			continue
		}

		modelName := strcase.ToCamel(value.Node.Name)
		var method_parameter string

		buf.WriteString("\n// " + modelName + " data model for the " + value.Node.Name + " structure\n")
		buf.WriteString("type " + modelName + " struct {\n")

		for _, field := range value.Node.Entries {
			typeField := _build_type_field(field.Field)

			if field.Field.Name == "parameters" {
				buf.WriteString("\t" + strings.ToLower(field.Field.Name) + typeField + "\n")
				method_parameter = _build_method_parameter(juaz, modelName, field.Field.Type.Reference)
			} else {
				typeField += " `" + `json:"` + strcase.ToSnake(field.Field.Name) + `,omitempty"` + "`"
				buf.WriteString("\t" + strcase.ToCamel(field.Field.Name) + typeField + "\n")
			}
		}

		buf.WriteString("}\n")
		buf.WriteString(method_parameter)
	}

	return buf.String()
}

func _build_type_field(field *grammar.Field) (t string) {
	if field.Type.Map != nil {
		t = "map[" + field.Type.Map.Key.Scalar.GoString() + "]" + field.Type.Map.Value.Scalar.GoString()
	} else if field.Type.Scalar != 0 {
		t = "*" + field.Type.Scalar.GoString()
	} else {
		t = "*" + strcase.ToCamel(field.Type.Reference)
	}

	if field.Repeated {
		t = " []" + t
	} else {
		t = " " + t
	}

	return t
}
