package proces

import (
	"regexp"
	"strconv"
	"strings"
)

func ReplaceNums(input string) string {
	// Преобразование бинарных и шестнадцатиричных чисел в десятичные
	for {
		converted := false
		input = regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).ReplaceAllStringFunc(input, func(match string) string {
			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*bin\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
			}
			binaryStr := parts[1]
			decimal, err := strconv.ParseInt(binaryStr, 2, 64)
			if err != nil {
				return match
			}
			converted = true
			return strconv.FormatInt(decimal, 10)
		})

		input = regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).ReplaceAllStringFunc(input, func(match string) string {
			parts := regexp.MustCompile(`\b(\w+)\s*\(\s*hex\s*\)`).FindStringSubmatch(match)
			if len(parts) < 2 {
				return match
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
	input = regexp.MustCompile(`\(\s*(bin|hex)\s*\)`).ReplaceAllString(input, "")
	return input
}

func AddSpace(str string) string {
	var newStr strings.Builder
	for i := 0; i < len(str); i++ {
		// Добавляем пробел перед открывающей скобкой, если его нет
		if str[i] == '(' && i > 0 {
			if str[i-1] != ' ' {
				newStr.WriteByte(' ')
			}
		}

		// Пишем текущий символ в новый строковый билд
		newStr.WriteByte(str[i])

		// Добавляем пробел после закрывающей скобки, если его нет
		if str[i] == ')' && i < len(str)-1 {
			if str[i+1] != ' ' {
				newStr.WriteByte(' ')
			}
		}
	}
	return newStr.String()
}
