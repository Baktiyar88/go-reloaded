package proces

import (
	"regexp"

	"strings"
)

func ProscessPuncQ(input string) string {
	input = FormatPunc(input)
	input = FixQuotes(input)

	return input
}

func FormatPunc(input string) string {
	re := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	input = re.ReplaceAllString(input, "$1")

	re = regexp.MustCompile(`([.,!?:;])([a-zA-Z0-9-])`)
	input = re.ReplaceAllString(input, "$1 $2")

	input = strings.Join(strings.Fields(input), " ")

	return strings.TrimSpace(input)

}

func FixQuotes(input string) string {
	input = fixDoubleQuotes(input)
	input = fixSingleQuotes(input)

	return input
}

func fixDoubleQuotes(input string) string {
	re := regexp.MustCompile(`"\s*(.*?)\s*"`)
	input = re.ReplaceAllString(input, `"$1"`)

	re = regexp.MustCompile(`(["])\s+([\'\w])`)
	input = re.ReplaceAllString(input, `$1$2`)
	re = regexp.MustCompile(`([\'\w])\s+(["])`)
	input = re.ReplaceAllString(input, `$1$2`)

	quoteCount := strings.Count(input, `"`)
	var result []rune
	inQuote := false

	for i := 0; i < len(input); i++ {
		currnetChar := rune(input[i])

		if currnetChar == '"' {
			if quoteCount%2 != 0 && strings.Count(string(result), `"`) == quoteCount-1 {
				result = append(result, currnetChar, ' ')
				continue
			}

			if !inQuote {
				inQuote = true
				if i > 0 && input[i-1] != ' ' && input[i-1] != '"' && input[i-1] != '\'' {
					result = append(result, ' ')

				}
				result = append(result, currnetChar)
			} else {
				inQuote = false
				result = append(result, currnetChar)

				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')

				}
			}
		} else {
			result = append(result, currnetChar)

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

	validSuffixes := map[string]bool{"t": true, "ll": true, "ve": true, "m": true, "s": true, "d": true, "re": true}

	quoteCount := strings.Count(input, "'")
	var result []rune
	inQuote := false

	for i := 0; i < len(input); i++ {
		currentChar := rune(input[i])
		if currentChar == '\'' {

			if i > 0 && i+1 < len(input) {

				prevWordEnd := i - 1
				for prevWordEnd >= 0 && isLetter(rune(input[prevWordEnd])) {
					prevWordEnd--
				}
				prevWord := input[prevWordEnd+1 : i]

				nextWordStart := i + 1
				for nextWordStart < len(input) && isLetter(rune(input[nextWordStart])) {
					nextWordStart++
				}
				nextSuffix := input[i+1 : nextWordStart]

				if isWord(prevWord) && validSuffixes[nextSuffix] {
					result = append(result, currentChar)
					continue
				}
			}

			if quoteCount == 1 {
				result = append(result, currentChar)

				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			if quoteCount%2 != 0 && strings.Count(string(result), `'`) == quoteCount-1 && !inQuote {
				result = append(result, currentChar)

				if i+1 < len(input) && !strings.ContainsAny(string(input[i+1]), ` .,;!?`) {
					result = append(result, ' ')
				}
				continue
			}

			if !inQuote {
				inQuote = true

				if i > 0 && input[i-1] != ' ' && input[i-1] != '\'' && input[i-1] != '"' {
					result = append(result, ' ')
				}
				result = append(result, currentChar)
			} else {

				inQuote = false
				result = append(result, currentChar)

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

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isWord(s string) bool {
	for _, ch := range s {
		if !isLetter(ch) {
			return false
		}
	}
	return len(s) > 0
}
