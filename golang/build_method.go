package golang

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_interface_methods(juaz *grammar.Juaz) string {
	pkg := strcase.ToCamel(juaz.Entries[0].Package)

	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("\n// I%sClient defines the interface of the provided methods\n", pkg))
	buf.WriteString(fmt.Sprintf("type I%sClient interface {\n", pkg))
	for _, value := range juaz.Entries {
		if value.Impl != nil {
			reference := strcase.ToCamel(value.Impl.Output.Reference)
			if value.Impl.Repeated {
				reference = fmt.Sprintf("[]%s", reference)
			}

			buf.WriteString(fmt.Sprintf("\t%s(context.Context, *%s) (*%s, error)\n", strcase.ToCamel(value.Impl.Name),
				strcase.ToCamel(value.Impl.Input.Reference), reference))
		}
	}
	buf.WriteString("}\n")
	return buf.String()
}
