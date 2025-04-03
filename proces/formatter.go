package proces

import (
	"regexp"
	"strings"
	"unicode"
)

func Articl(input string) string {
	// Список исключений, где артикль не меняется
	exceptions := map[string]bool{
		"for":  true,
		"and":  true,
		"or":   true,
		"yet":  true,
		"with": true,
		"so":   true,
		"but":  true,
	}

	// Регулярное выражение для слов с немым "h" (например, "hour", "honor")
	silentHRegex := regexp.MustCompile(`(?i)^h(onor|our|heir|hour)`)

	// Разделяем строку на слова, учитывая кавычки
	words := splitWithQuotes(input)

	// Проходим по словам, проверяя артикли
	for i := 0; i < len(words)-1; i++ {
		if isArticle(words[i]) { // Если текущее слово — артикль
			nextWord := words[i+1]
			// Извлекаем слово без кавычек для проверки
			nextWordClean := strings.Trim(nextWord, "'")
			nextWordLower := strings.ToLower(nextWordClean)

			// Если следующее слово — артикль или исключение, оставляем текущий артикль без изменений
			if isArticle(nextWordClean) || exceptions[nextWordLower] {
				continue
			}

			// Если следующее слово не артикль и не исключение, корректируем артикль
			if len(nextWordClean) > 0 {
				firstLetter := strings.ToLower(string(nextWordClean[0]))
				isVowel := strings.Contains("aeiou", firstLetter)
				isSilentH := silentHRegex.MatchString(nextWordClean)
				correctArticle := "a"
				if isVowel || isSilentH {
					correctArticle = "an"
				}

				// Учитываем регистр исходного артикля
				if isUpper(words[i]) && isUpper(nextWordClean) {
					words[i] = strings.ToUpper(correctArticle) // "A" или "AN"
				} else if unicode.IsUpper(rune(words[i][0])) {
					words[i] = strings.Title(correctArticle) // "A" или "An"
				} else {
					words[i] = correctArticle // "a" или "an"
				}
			}
		}
	}

	// Собираем строку обратно
	return strings.Join(words, " ")
}

// isArticle проверяет, является ли слово артиклем
func isArticle(word string) bool {
	lower := strings.ToLower(word)
	return lower == "a" || lower == "an"
}

// isUpper проверяет, написано ли слово заглавными буквами
func isUpper(word string) bool {
	return word == strings.ToUpper(word)
}

// splitWithQuotes разбивает строку, сохраняя слова в кавычках
func splitWithQuotes(input string) []string {
	re := regexp.MustCompile(`'[^']*'|\S+`)
	return re.FindAllString(input, -1)
}
