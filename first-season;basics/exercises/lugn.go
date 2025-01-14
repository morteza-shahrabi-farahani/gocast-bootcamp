package exercises

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	sum := 0

	if len(id) <= 1 {
		return false
	}

	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	double := false
	for i := len(id) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if double {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}

		sum += digit
		double = !double
	}

	if sum%10 == 0 || sum < 10 {
		return true
	} else {
		return false
	}
}
