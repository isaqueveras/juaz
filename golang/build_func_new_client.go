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

func _build_func_new_client(juaz *grammar.Juaz) string {
	var buf strings.Builder
	pkg := strcase.ToCamel(juaz.Entries[0].Package)
	buf.WriteString("\nfunc New" + strcase.ToCamel(pkg) + "Client(cc juazeiro.ClientConnInterface) I" + strcase.ToCamel(pkg) + "Client {\n")
	buf.WriteString("\treturn &" + strcase.ToCamel(pkg) + "Client{cc: cc}\n")
	buf.WriteString("}\n")
	return buf.String()
}
