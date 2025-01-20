package exercises

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	results := []string{}
	for i := 0; i < len(rhyme)-1; i++ {
		results = append(results, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	results = append(results, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return results
}
