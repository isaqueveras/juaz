package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_struct_client(juaz *grammar.Juaz) string {
	var buf strings.Builder
	pkg := strcase.ToCamel(juaz.Entries[0].Package)
	buf.WriteString("\ntype " + strcase.ToCamel(pkg) + "Client struct {\n")
	buf.WriteString("\tcc juazeiro.ClientConnInterface\n")
	buf.WriteString("}\n")
	return buf.String()
}
