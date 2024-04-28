//  Created by Isaque Veras on 03/15/24.
//  Copyright © 2024 Isaque Veras. All rights reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package golang

import (
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/isaqueveras/juaz/grammar"
)

type paramsModelOut struct {
	name     string
	buf      string
	nameFunc string
}

func _build_parameters(juaz *grammar.Juaz) (out []paramsModelOut) {
	for _, entry := range juaz.Entries {
		if entry.Node == nil {
			continue
		}

		for _, value := range entry.Node.Entries {
			if value.Field.Name != "parameters" {
				continue
			}

			for _, entry2 := range juaz.Entries {
				if entry2.Node == nil {
					continue
				}

				if value.Field.Type.Reference == entry2.Node.Name {
					nameFuncInSnake := strcase.ToSnake(value.Field.Type.Reference)

					for i := range out {
						if out[i].nameFunc == nameFuncInSnake {
							return
						}
					}

					var buf strings.Builder
					buf.WriteString("\nfunc _build_" + nameFuncInSnake + "_parameters(in *" + strcase.ToCamel(value.Field.Type.Reference) + ") string {\n")
					buf.WriteString("\tval := &url.Values{}\n")

					var hasImportNetUrl bool
					for idx := range imports {
						if imports[idx] == "net/url" {
							hasImportNetUrl = true
						}
					}

					if !hasImportNetUrl {
						imports = append(imports, "net/url")
					}

					for i := range entry2.Node.Entries {
						if entry2.Node.Entries[i].Field.Repeated {
							buf.WriteString("\tfor _, value := range in." + strcase.ToCamel(entry2.Node.Entries[i].Field.Name) + " {\n")
							buf.WriteString("\t\tif value == nil {\n")
							buf.WriteString("\t\t\tcontinue\n")
							buf.WriteString("\t\t}\n")

							if entry2.Node.Entries[i].Field.Type.Reference != "" {
								buf.WriteString("\t\t" + `val.Add("` + strcase.ToSnake(entry2.Node.Entries[i].Field.Name) + `", value.String())` + "\n")
							} else {
								buf.WriteString("\t\t" + `val.Add("` + strcase.ToSnake(entry2.Node.Entries[i].Field.Name) + `", fmt.Sprintf("%v", *value))` + "\n")
							}
							buf.WriteString("\t}\n")
						} else {
							buf.WriteString("\tif in." + strcase.ToCamel(entry2.Node.Entries[i].Field.Name) + " != nil {\n")
							if entry2.Node.Entries[i].Field.Type.Reference != "" {
								buf.WriteString("\t\t" + `val.Add("` + strcase.ToSnake(entry2.Node.Entries[i].Field.Name) + `", fmt.Sprintf("%v", in.` + strcase.ToCamel(entry2.Node.Entries[i].Field.Name) + `.String()))` + "\n")
							} else {
								buf.WriteString("\t\t" + `val.Add("` + strcase.ToSnake(entry2.Node.Entries[i].Field.Name) + `", fmt.Sprintf("%v", *in.` + strcase.ToCamel(entry2.Node.Entries[i].Field.Name) + `))` + "\n")
							}
							buf.WriteString("\t}\n")
						}
					}

					buf.WriteString("\t" + `return ` + `"?" + ` + `val.Encode()` + "\n")
					buf.WriteString("}\n")

					out = append(out, paramsModelOut{name: entry.Node.Name, buf: buf.String(), nameFunc: nameFuncInSnake})
				}
			}
		}
	}

	return out
}
