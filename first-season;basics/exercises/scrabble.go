package exercises

import "strings"

func Score(word string) int {
	lowerWord := strings.ToLower(word)
	result := 0

	for _, char := range lowerWord {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'l' ||
			char == 'n' || char == 'r' || char == 's' || char == 't' {
			result += 1
		} else if char == 'd' || char == 'g' {
			result += 2
		} else if char == 'b' || char == 'c' || char == 'm' || char == 'p' {
			result += 3
		} else if char == 'f' || char == 'h' || char == 'v' || char == 'w' || char == 'y' {
			result += 4
		} else if char == 'k' {
			result += 5
		} else if char == 'j' || char == 'x' {
			result += 8
		} else if char == 'q' || char == 'z' {
			result += 10
		}
	}

	return result
}
