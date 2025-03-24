package words

import (
	"strings"
)

// ToCamelCase converts a string to camelCase.
func ToCamelCase(s string) string {
	words := strings.Fields(s)
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}
