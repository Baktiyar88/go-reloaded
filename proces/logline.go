package proces

import (
	"regexp"
	"strconv"
	"strings"
)

// ProcessLine обрабатывает строку, применяя все необходимые преобразования
func ProcessLine(line string) string {
	tokens := Tokenize(line)
	tokens = ProcessTokens(tokens)
	output := Assemble(tokens)
	output = FixPunctuationSpacing(output)
	output = HandleQuotes(output)
	output = HandleAAn(output)
	return output
}

// Tokenize разбивает строку на токены
func Tokenize(text string) []string {
	re := regexp.MustCompile(`(\(hex\)|\(bin\)|\(up\)|\(low\)|\(cap\)|\(up,\s*\d+\)|\(low,\s*\d+\)|\(cap,\s*\d+\)|\.\.\.|['"]|[.,!?;:]+|\w+|[\S]+)`)
	return re.FindAllString(text, -1)
}

// ProcessTokens обрабатывает токены, применяя команды
func ProcessTokens(tokens []string) []string {
	var result []string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
			marker := strings.Trim(token, "()")
			parts := strings.Split(marker, ",")
			cmd := strings.TrimSpace(parts[0])
			n := 1

			if len(parts) > 1 {
				var err error
				n, err = strconv.Atoi(strings.TrimSpace(parts[1]))
				if err != nil {
					n = 1
				}
			}

			words := findLastNWords(result, n)
			for _, idx := range words {
				switch cmd {
				case "hex":
					if num, err := strconv.ParseInt(result[idx], 16, 64); err == nil {
						result[idx] = strconv.FormatInt(num, 10)
					}
				case "bin":
					if num, err := strconv.ParseInt(result[idx], 2, 64); err == nil {
						result[idx] = strconv.FormatInt(num, 10)
					}
				case "up":
					result[idx] = strings.ToUpper(result[idx])
				case "low":
					result[idx] = strings.ToLower(result[idx])
				case "cap":
					result[idx] = ToTitleCase(strings.ToLower(result[idx]))
				}
			}
		} else {
			result = append(result, token)
		}
	}
	return result
}

// Assemble собирает токены в строку
func Assemble(tokens []string) string {
	var b strings.Builder
	for i, t := range tokens {
		if i > 0 && isWord(t) {
			b.WriteString(" ")
		}
		b.WriteString(t)
	}
	return b.String()
}

// findLastNWords находит последние N слов в результате
func findLastNWords(result []string, n int) []int {
	indices := []int{}
	for i := len(result) - 1; i >= 0 && len(indices) < n; i-- {
		if isWord(result[i]) {
			indices = append([]int{i}, indices...)
		}
	}
	return indices
}

// isWord проверяет, является ли токен словом
func isWord(token string) bool {
	for _, c := range token {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			return true
		}
	}
	return false
}
