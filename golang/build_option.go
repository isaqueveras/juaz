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

import "github.com/isaqueveras/juaz/grammar"

func _build_status_option(juaz *grammar.Juaz) string {
	for _, value := range juaz.Entries {
		if value.Option != nil && value.Option.Name == "status" {
			return value.Option.Value
		}
	}
	return ""
}

func _build_method_option(juaz *grammar.Juaz) string {
	for _, value := range juaz.Entries {
		if value.Option != nil && value.Option.Name == "method" {
			return value.Option.Value
		}
	}
	return ""
}
