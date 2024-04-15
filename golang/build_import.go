package golang

import "strings"

const importLibJuazeiro = `"github.com/isaqueveras/juazeiro"`

func _build_import() string {
	var buf strings.Builder
	buf.WriteString("import (\n")
	for _, value := range imports {
		buf.WriteString("\t" + `"` + value + `"` + "\n")
	}
	buf.WriteString("\n\t" + importLibJuazeiro + "\n")
	buf.WriteString(")\n")
	return buf.String()
}
