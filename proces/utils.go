package proces

import (
	"unicode"
)

// ToTitleCase преобразует строку в заглавный регистр
func ToTitleCase(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToTitle(runes[0])
	}
	return string(runes)
}
