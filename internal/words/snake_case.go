package words

import "unicode"

// ToSnakeCase converts a string to snake_case.
func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else if r == ' ' {
			result = append(result, '_')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
