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
			if value.Impl.Output != nil {
				reference := strcase.ToCamel(value.Impl.Output.Reference)
				if value.Impl.Output.Scalar != 0 {
					reference = value.Impl.Output.Scalar.GoString()
				}

				if value.Impl.Repeated {
					reference = fmt.Sprintf("[]%s", reference)
				}

				buf.WriteString(fmt.Sprintf("\t%s(context.Context, *%s) (*%s, error)\n", strcase.ToCamel(value.Impl.Name),
					strcase.ToCamel(value.Impl.Input.Reference), reference))
			} else {
				buf.WriteString(fmt.Sprintf("\t%s(context.Context, *%s) error\n", strcase.ToCamel(value.Impl.Name),
					strcase.ToCamel(value.Impl.Input.Reference)))
			}
		}
	}
	buf.WriteString("}\n")
	return buf.String()
}
