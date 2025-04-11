package proces

import (
	"regexp"
	"strings"
)

// ProscessPuncQ обрабатывает входную строку, форматируя пунктуацию и кавычки
func ProscessPuncQ(input string) string {
	input = FormatPunc(input)
	input = FixQuotes(input)

	return input
}

func FormatPunc(input string) string {
	// Удаляем лишние пробелы вокруг знаков препинания
	re := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	input = re.ReplaceAllString(input, "$1")

	// Добавляем пробел после знаков препинания, если за ними идет буква или цифра
	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")

	// Удаляем множественные пробелы и объединяем слова
	input = strings.Join(strings.Fields(input), " ")

	return strings.TrimSpace(input)
}

// FixQuotes исправляет форматирование одинарных и двойных кавычек
func FixQuotes(input string) string {
	input = fixDoubleQuotes(input)
	input = fixSingleQuotes(input)

	return input
}

func fixDoubleQuotes(input string) string {
	// Удаляем пробелы внутри содержимого в двойных кавычках
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)

	// Удаляем пробелы между кавычкой и следующим символом
	re = regexp.MustCompile(`(["])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	// Удаляем пробелы между символом и следующей кавычкой
	re = regexp.MustCompile(`([\'\w])\s+(["])`)
	input = re.ReplaceAllString(input, `$1$2`)

	quoteCount := strings.Count(input, `"`) // Подсчитываем количество кавычек
	var result []rune                       // Создаем срез для результата
	inQuote := false

	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])

		if currentChar == '"' {
			// Обрабатываем последнюю кавычку для нечетного количества
			if quoteCount%2 != 0 && strings.Count(string(result), `"`) == quoteCount-1 {
				result = append(result, currentChar, ' ')
				continue
			}

			if !inQuote {
				inQuote = true
				// Добавляем пробел перед кавычкой, если перед ней нет пробела или другой кавычки
				if i > 0 && input[i-1] != ' ' && input[i-1] != '"' && input[i-1] != '\'' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				inQuote = false
				result = append(result, currentChar)
				// Добавляем пробел после кавычки, если за ней нет пунктуации
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, currentChar)
		}
	}

	return strings.TrimSpace(string(result))
}

func fixSingleQuotes(input string) string {

	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	input = re.ReplaceAllString(input, "'$1'")

	re = regexp.MustCompile(`(['])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)

	re = regexp.MustCompile(`([\'\w])\s+(['])`)
	input = re.ReplaceAllString(input, `$1$2`)

	// Определяем допустимые суффиксы для сокращений (например, 't, 'll, 's)
	validSuffixes := map[string]bool{"t": true, "ll": true, "ve": true, "m": true, "s": true, "d": true, "re": true}

	quoteCount := strings.Count(input, "'") // Подсчитываем количество одинарных кавычек
	var result []rune                       // Создаем срез для результата
	inQuote := false

	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '\'' {
			// Проверяем, является ли кавычка частью сокращения
			if i > 0 && i+1 < len(input) {

				prevWordEnd := i - 1
				for prevWordEnd >= 0 && isLetter(rune(input[prevWordEnd])) {
					prevWordEnd--
				}
				prevWord := input[prevWordEnd+1 : i]

				// Находим следующий суффикс
				nextWordStart := i + 1
				for nextWordStart < len(input) && isLetter(rune(input[nextWordStart])) {
					nextWordStart++
				}
				nextSuffix := input[i+1 : nextWordStart]

				// Если это сокращение, добавляем кавычку без изменений
				if isWord(prevWord) && validSuffixes[nextSuffix] {
					result = append(result, currentChar)
					continue
				}
			}

			// // Обрабатываем одиночную кавычку
			if quoteCount == 1 {
				result = append(result, currentChar)
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			// Обрабатываем последнюю кавычку для нечетного количества
			if quoteCount%2 != 0 && strings.Count(string(result), `'`) == quoteCount-1 && !inQuote {
				result = append(result, currentChar)
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			if !inQuote {
				inQuote = true
				// Добавляем пробел перед кавычкой, если перед ней нет пробела или другой кавычки
				if i > 0 && input[i-1] != ' ' && input[i-1] != '\'' && input[i-1] != '"' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {
				inQuote = false
				result = append(result, currentChar)
				// Добавляем пробел после кавычки, если за ней нет пунктуации
				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
			}
		} else {
			result = append(result, currentChar)
		}
	}

	return strings.TrimSpace(string(result))
}

// isLetter проверяет, является ли символ буквой
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isWord проверяет, является ли строка словом (содержит только буквы)
func isWord(s string) bool {
	for _, ch := range s {
		if !isLetter(ch) {
			return false
		}
	}
	return len(s) > 0
}
