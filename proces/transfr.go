package proces

import (
	"fmt"
	"regexp"
	"strings"
)

func ProcessTrans(inputText string) string {
	// Define a regex to find modification tags like (up), (low, 2), etc.
	modsRegex := regexp.MustCompile(`\((\s*up\s*|\s*low\s*|\s*cap\s*)\s*(\s*,\s*\d+\s*)?\)`)
	for {
		changed := false
		match := modsRegex.FindStringIndex(inputText)
		if match != nil {
			prefix := strings.TrimSpace(inputText[:match[0]])
			suffix := inputText[match[1]:]
			matchMod := inputText[match[0]:match[1]]
			mod, count := getModAndCount(matchMod)
			if !hasLetterOrDigit(prefix) {
				// If there are no valid words before the tag, skip this tag
				inputText = prefix + suffix
				continue
			}
			// Apply the modification to the prefix
			modifiedPrefix := modificatePrefix(mod, count, prefix)
			// Update the input text
			inputText = modifiedPrefix + suffix
			changed = true
		}
		if !changed {
			break
		}
	}

	return strings.TrimSpace(inputText)
}

// getModAndCount parses the modification tag to extract the action and count.
func getModAndCount(tag string) (string, int) {
	parts := strings.Split(strings.Trim(tag, "()"), ",")
	mod := strings.TrimSpace(parts[0])
	count := 1
	if len(parts) > 1 {
		fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &count)
	}
	return mod, count
}

// modificatePrefix applies the specified modification to the prefix.
func modificatePrefix(mod string, count int, prefix string) string {
	words := strings.Fields(prefix)
	lastWordIndex := len(words) - 1
	switch mod {
	case "up":
		for i := 0; i < count && lastWordIndex-i >= 0; i++ {
			words[lastWordIndex-i] = strings.ToUpper(words[lastWordIndex-i])
		}
	case "low":
		for i := 0; i < count && lastWordIndex-i >= 0; i++ {
			words[lastWordIndex-i] = strings.ToLower(words[lastWordIndex-i])
		}
	case "cap":
		for i := 0; i < count && lastWordIndex-i >= 0; i++ {
			words[lastWordIndex-i] = customTitle(words[lastWordIndex-i])
		}
	}
	return strings.Join(words, " ")
}

// customTitle capitalizes the first letter of a word and lowers the rest.
func customTitle(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	}
	return s
}

// hasLetterOrDigit checks if the input string contains at least one letter or digit.
func hasLetterOrDigit(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}
