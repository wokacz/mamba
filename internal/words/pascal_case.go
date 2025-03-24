package words

import (
	"strings"
	"unicode"
)

// ToPascalCase converts a string to PascalCase.
func ToPascalCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == ' ' || unicode.IsUpper(r)
	})
	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, "")
}
