package proces

import (
	"regexp"
	"strconv"
	"strings"
)

// ReplaceNums преобразует бинарные и шестнадцатеричные числа в десятичные, удаляя соответствующие теги
func ReplaceNums(input string) string {

	for {
		converted := false // Флаг для отслеживания успешных преобразований

		input = regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).ReplaceAllStringFunc(input, func(match string) string {
			// Извлекаем бинарное число из совпадения
			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match // Возвращаем без изменений, если формат неверный
			}
			binaryStr := parts[1] // Получаем строку с бинарным числом
			// Преобразуем бинарное число в десятичное
			decimal, err := strconv.ParseInt(binaryStr, 2, 64)
			if err != nil {
				return match
			}
			converted = true

			return strconv.FormatInt(decimal, 10)
		})

		// Обработка шестнадцатеричных чисел, помеченных как (hex)
		input = regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).ReplaceAllStringFunc(input, func(match string) string {

			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match // Возвращаем без изменений, если формат неверный
			}
			hexStr := parts[1]

			decimal, err := strconv.ParseInt(hexStr, 16, 64)
			if err != nil {
				return match
			}
			converted = true

			return strconv.FormatInt(decimal, 10)
		})

		if !converted {
			break
		}
	}
	// Удаляем оставшиеся теги (bin) или (hex) без чисел
	input = regexp.MustCompile(`\(\s*(bin|hex)\s*\)`).ReplaceAllString(input, "")

	return input
}

// AddSpace добавляет пробелы вокруг скобок для улучшения форматирования
func AddSpace(str string) string {
	var newStr strings.Builder
	for i := 0; i < len(str); i++ {
		// Добавляем пробел перед открывающей скобкой, если его нет
		if str[i] == '(' && i > 0 {
			if str[i-1] != ' ' {
				newStr.WriteByte(' ')
			}
		}

		// Записываем текущий символ
		newStr.WriteByte(str[i])

		if str[i] == ')' && i < len(str)-1 {
			if str[i+1] != ' ' {
				newStr.WriteByte(' ')
			}
		}
	}

	return newStr.String()
}
