package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_type(juaz *grammar.Juaz) string {
	var buf strings.Builder

	for _, value := range juaz.Entries {
		if value.Type != nil {
			enumName := strcase.ToCamel(*value.Type.Name)
			buf.WriteString("\n// " + enumName + " defines the " + strings.ToLower(enumName) + " enum type\n")
			buf.WriteString("type " + enumName + " string\n\n")

			buf.WriteString("const (\n")
			for _, enum := range value.Type.Values {
				buf.WriteString("\t" + enumName + strcase.ToCamel(enum.Value.Key) + "\t" + enumName + " = " + `"` + enum.Value.Key + `"` + "\n")
			}
			buf.WriteString(")\n")

			buf.WriteString("\n// String convert " + strings.ToLower(enumName) + " type to string\n")
			buf.WriteString("func (t " + enumName + ") String() string {\n")
			buf.WriteString("\treturn string(t)\n")
			buf.WriteString("}\n")
		}
	}

	return buf.String()
}
