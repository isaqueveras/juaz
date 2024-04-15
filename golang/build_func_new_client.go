package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_func_new_client(juaz *grammar.Juaz) string {
	var buf strings.Builder
	pkg := strcase.ToCamel(juaz.Entries[0].Package)
	buf.WriteString("\nfunc New" + strcase.ToCamel(pkg) + "Client(cc juazeiro.ClientConnInterface) I" + strcase.ToCamel(pkg) + "Client {\n")
	buf.WriteString("\treturn &" + strcase.ToCamel(pkg) + "Client{cc: cc}\n")
	buf.WriteString("}\n")
	return buf.String()
}
