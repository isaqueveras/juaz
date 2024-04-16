package golang

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

func _build_implement_method(juaz *grammar.Juaz) string {
	var buf strings.Builder
	pkg := strcase.ToCamel(juaz.Entries[0].Package)
	for _, value := range juaz.Entries {
		if value.Impl != nil {
			method := strcase.ToCamel(value.Impl.Name) + `(ctx context.Context, in *` + strcase.ToCamel(value.Impl.Input.Reference) + `) (*` + strcase.ToCamel(value.Impl.Output.Reference) + `, error)`
			buf.WriteString("\n// " + strcase.ToCamel(value.Impl.Name) + " implements the " + strcase.ToCamel(value.Impl.Name) + " method of the interface\n")
			buf.WriteString("func (c *" + strcase.ToCamel(pkg) + "Client) " + method + " {\n")
			buf.WriteString("\tout := new(" + strcase.ToCamel(value.Impl.Output.Reference) + ")\n")

			var httpMethod, uri, status string
			for _, v := range value.Impl.Entry {
				switch v.Name {
				case "uri":
					var path string
					var args []string
					for _, value := range strings.Split(v.Value, "/") {
						if strings.Contains(value, "{") {
							args = append(args, value)
							value = strings.Replace(value, value, "%v", -1)
						}
						path += value + "/"
					}

					if len(args) > 0 {
						path = strings.TrimSuffix(path, "/")
						path = strings.TrimLeft(path, `"`)
						path = strings.TrimRight(path, `"`)

						var arg string
						for i := range args {
							args[i] = strings.TrimLeft(args[i], `{`)
							args[i] = strings.TrimRight(args[i], `}`)
							arg += " *in." + strcase.ToCamel(args[i]) + ","
						}

						uri = `fmt.Sprintf("` + path + `",` + strings.TrimRight(arg, `,`) + ")"
					} else {
						uri = v.Value
					}
				case "method":
					httpMethod = _build_http_method(v.Value)
				case "status":
					status = _build_status_code(v.Value)
				}
			}

			var hasImportFmt bool
			for idx := range imports {
				if imports[idx] == "fmt" {
					hasImportFmt = true
				}
			}

			if !hasImportFmt {
				imports = append(imports, "fmt")
			}

			buf.WriteString("\t" + `uri := ` + uri + "\n")

			for _, params := range _build_parameters(juaz) {
				if params.name == value.Impl.Input.Reference {
					buf.WriteString("\tif in.parameters != nil {\n")
					buf.WriteString(fmt.Sprintf("\t\turi += _build_%s_parameters(in.parameters)\n", params.nameFunc))
					buf.WriteString("\t\tin.parameters = nil\n")
					buf.WriteString("\t}\n")
				}
			}

			buf.WriteString(fmt.Sprintf("\terr := c.cc.Invoke(ctx, %s, uri, %s, in, out)\n", httpMethod, status))
			buf.WriteString("\treturn out, err\n")
			buf.WriteString("}\n")
		}
	}

	return buf.String()
}

func _build_http_method(method string) string {
	switch method {
	case `"GET"`:
		return "http.MethodGet"
	case `"POST"`:
		return "http.MethodPost"
	case `"PUT"`:
		return "http.MethodPut"
	case `"DELETE"`:
		return "http.MethodDelete"
	default:
		panic("method http not supported")
	}
}
