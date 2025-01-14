package exercises

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	result := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("length must be equal")
	}

	for index, _ := range a {
		if a[index] != b[index] {
			result += 1
		}
	}

	return result, nil
}
