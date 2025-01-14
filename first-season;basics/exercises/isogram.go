package exercises

import "strings"

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)

	wordsMap := make(map[rune]bool)

	for _, alphabet := range lowerWord {
		if alphabet == ' ' || alphabet == '-' {
			continue
		}

		if wordsMap[alphabet] {
			return false
		} else {
			wordsMap[alphabet] = true
		}
	}

	return true
}
