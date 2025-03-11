package goreload

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	text, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	tokens := tokenize(string(text))
	result := processTokens(tokens)
	output := assemble(result)
	output = handleQuotes(output)
	output = handleAAn(output)
	err = os.WriteFile(os.Args[2], []byte(output), 0o644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}
}

func tokenize(text string) []string {
	re := regexp.MustCompile(`(\w+|\(hex\)|\(bin\)|\(up\)|\(low\)|\(cap\)|\(up,\s*\d+\)|\(low,\s*\d+\)|\(cap,\s*\d+\)|[.,!?;:]+|\.\.\.|\S+)`)
	return re.FindAllString(text, -1)
}

func processTokens(tokens []string) []string {
	var result []string
	for i := 0; i < len(tokens); i++ {
		if strings.HasPrefix(tokens[i], "(") && strings.HasSuffix(tokens[i], ")") {
			marker := strings.Trim(tokens[i], "()")
			parts := strings.Split(marker, ",")
			cmd := strings.TrimSpace(parts[0])
			n := 1
			if len(parts) > 1 {
				n, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
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
					result[idx] = strings.Title(strings.ToLower(result[idx]))
				}
			}
		} else {
			result = append(result, tokens[i])
		}
	}
	return result
}

func findLastNWords(result []string, n int) []int {
	indices := []int{}
	for i := len(result) - 1; i >= 0 && len(indices) < n; i-- {
		if isWord(result[i]) {
			indices = append([]int{i}, indices...)
		}
	}
	return indices
}

func isWord(token string) bool {
	for _, c := range token {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func assemble(result []string) string {
	var b strings.Builder
	for i, t := range result {
		if i > 0 && isWord(t) && !strings.ContainsAny(result[i-1], ".,!?;:") {
			b.WriteString(" ")
		}
		b.WriteString(t)
	}
	return b.String()
}

func handleQuotes(output string) string {
	return regexp.MustCompile(`'\s*(.*?)\s*'`).ReplaceAllString(output, "'$1'")
}

func handleAAn(output string) string {
	return regexp.MustCompile(`\b(a|A)\s+([aeiouhAEIOUH]\w*)`).ReplaceAllStringFunc(output, func(m string) string {
		parts := strings.Fields(m)
		if strings.ToLower(parts[0]) == "a" {
			return "an " + parts[1]
		}
		return "An " + parts[1]
	})
}
