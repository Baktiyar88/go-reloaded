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

// isVowelSound проверяет, начинается ли слово с гласного звука
func isVowelSound(word string) bool {
	lowerWord := strings.ToLower(word)

	if strings.HasPrefix(lowerWord, "a") ||
		strings.HasPrefix(lowerWord, "e") ||
		strings.HasPrefix(lowerWord, "i") ||
		strings.HasPrefix(lowerWord, "o") ||
		strings.HasPrefix(lowerWord, "u") {
		return true
	}

	if strings.HasPrefix(lowerWord, "hour") ||
		strings.HasPrefix(lowerWord, "heir") ||
		strings.HasPrefix(lowerWord, "honest") ||
		strings.HasPrefix(lowerWord, "honour") {
		return true
	}
	return false
}

// correctArticle возвращает правильный артикл в зависимости от слова и регистра
func correctArticle(article, word string) string {
	// Определяем, является ли первый символ артикля заглавным
	isUpper := strings.ToUpper(article[:1]) == article[:1]

	if isVowelSound(word) {
		if isUpper {
			return "An"
		}
		return "an"
	} else {
		if isUpper {
			return "A"
		}
		return "a"
	}
}

// HandleAAn корректирует артикли в тексте
func HandleAAn(text string) string {
	// Регулярное выражение для поиска артикля и следующего слова
	re := regexp.MustCompile(`\b(a|an|A|An|aN|AN)\s+(\w+)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		article := parts[0]
		word := parts[1]
		correctedArticle := correctArticle(article, word)
		return correctedArticle + " " + word
	})
}
