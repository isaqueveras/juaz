package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_interface_methods(juaz *grammar.Juaz) string {
	pkg := strcase.ToCamel(juaz.Entries[0].Package)

	var buf strings.Builder
	buf.WriteString("\n// I" + pkg + "Client defines the interface of the provided methods\n")
	buf.WriteString("type I" + pkg + "Client interface {\n")
	for _, value := range juaz.Entries {
		if value.Impl != nil {
			buf.WriteString("\t" + strcase.ToCamel(value.Impl.Name) + `(context.Context, *` + strcase.ToCamel(value.Impl.Input.Reference) + `) (*` + strcase.ToCamel(value.Impl.Output.Reference) + `, error)` + "\n")
		}
	}
	buf.WriteString("}\n")
	return buf.String()
}
