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
		if value.Impl == nil {
			continue
		}

		var method, out string
		if value.Impl.Output != nil {
			reference := strcase.ToCamel(value.Impl.Output.Reference)
			if value.Impl.Output.Scalar != 0 {
				reference = value.Impl.Output.Scalar.GoString()
			}

			if value.Impl.Repeated {
				reference = fmt.Sprintf("[]%s", reference)
			}

			method = fmt.Sprintf("%s(ctx context.Context, in *%s) (*%s, error)", strcase.ToCamel(value.Impl.Name),
				strcase.ToCamel(value.Impl.Input.Reference), reference)

			out = "\tout := new(" + reference + ")\n"
		} else {
			method = fmt.Sprintf("%s(ctx context.Context, in *%s) error", strcase.ToCamel(value.Impl.Name), strcase.ToCamel(value.Impl.Input.Reference))
		}

		buf.WriteString(fmt.Sprintf("\n// %s implements the %s method of the interface\n", strcase.ToCamel(value.Impl.Name), value.Impl.Name))
		buf.WriteString(fmt.Sprintf("func (c *%sClient) "+method+" {\n", strcase.ToCamel(pkg)))

		if out != "" {
			buf.WriteString(out)
		}

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
					uri = `"` + v.Value + `"`
				}
			case "method":
				httpMethod = _build_http_method(v.Value)
			case "status":
				status = _build_status_code(v.Value)
			}

			if status == "" {
				status = _build_status_code(_build_status_option(juaz))
			}

			if httpMethod == "" {
				httpMethod = _build_http_method(_build_method_option(juaz))
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

		if value.Impl.Output != nil {
			buf.WriteString(fmt.Sprintf("\terr := c.cc.Invoke(ctx, %s, uri, %s, in, out)\n", httpMethod, status))
			buf.WriteString("\treturn out, err\n")
		} else {
			buf.WriteString(fmt.Sprintf("\treturn c.cc.Invoke(ctx, %s, uri, %s, in, in)\n", httpMethod, status))
		}

		buf.WriteString("}\n")
	}

	return buf.String()
}

func _build_http_method(method string) string {
	switch method {
	case "GET":
		return "http.MethodGet"
	case "POST":
		return "http.MethodPost"
	case "PUT":
		return "http.MethodPut"
	case "DELETE":
		return "http.MethodDelete"
	default:
		panic("method http not supported")
	}
}
