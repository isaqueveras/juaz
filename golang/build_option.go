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
