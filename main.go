package main

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

	lines := strings.Split(string(text), "\n")
	var resultLines []string
	for _, line := range lines {
		tokens := tokenize(line)
		result := processTokens(tokens)
		output := assemble(result)
		output = fixPunctuationSpacing(output)
		output = handleQuotes(output)
		output = handleAAn(output)
		resultLines = append(resultLines, output)
	}

	finalOutput := strings.Join(resultLines, "\n")
	err = os.WriteFile(os.Args[2], []byte(finalOutput), 0o644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}
}

func tokenize(text string) []string {
	re := regexp.MustCompile(`(\w+|\(hex\)|\(bin\)|\(up\)|\(low\)|\(cap\)|\(up,\s*\d+\)|\(low,\s*\d+\)|\(cap,\s*\d+\)|\.\.\.|[.,!?;:]+|\S+)`)
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
					result[idx] = toTitleCase(strings.ToLower(result[idx]))
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
		if i > 0 && isWord(t) {
			b.WriteString(" ")
		}
		b.WriteString(t)
	}
	return b.String()
}

func fixPunctuationSpacing(text string) string {
	text = regexp.MustCompile(`\s*(\.\.\.)\s*`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`\s*([.,!?;:]+)\s*`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([.,!?;:])(\w)`).ReplaceAllString(text, "$1 $2")
	return text
}

func handleQuotes(output string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(output, "'$1'")
}

func handleAAn(output string) string {
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

// uefhrdiuberig
func toTitleCase(s string) string {
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToTitle(runes[0])
	}
	return string(runes)
}
