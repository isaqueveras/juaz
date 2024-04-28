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
