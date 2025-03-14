package proces

import (
	"regexp"
	"strings"
)

// FixPunctuationSpacing исправляет пробелы вокруг пунктуации
func FixPunctuationSpacing(text string) string {
	text = regexp.MustCompile(`\s*(\.\.\.)\s*`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`\s*([.,!?;:]+)\s*`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([.,!?;:])(\w)`).ReplaceAllString(text, "$1 $2")
	return text
}

// HandleQuotes обрабатывает кавычки
func HandleQuotes(output string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(output, "'$1'")
}

// HandleAAn исправляет артикли "a" и "an"
func HandleAAn(output string) string {
	re := regexp.MustCompile(`\b(a|A)\s+(\w+)`)
	return re.ReplaceAllStringFunc(output, func(m string) string {
		parts := strings.Fields(m)
		word := parts[1]
		if isVowelSound(word) {
			if parts[0] == "A" {
				return "An " + word
			}
			return "an " + word
		}
		return parts[0] + " " + word
	})
}

// isVowelSound проверяет, начинается ли слово с гласного звука
func isVowelSound(word string) bool {
	vowelSounds := []string{"a", "e", "i", "o", "u", "h"}
	lowerWord := strings.ToLower(word)
	for _, v := range vowelSounds {
		if strings.HasPrefix(lowerWord, v) && (v != "h" || lowerWord == "hour" || lowerWord == "heir" || lowerWord == "honest") {
			return true
		}
	}
	return false
}
