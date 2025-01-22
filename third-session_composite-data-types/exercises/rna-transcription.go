package exercises

func ToRNA(dna string) string {
	result := ""

	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			result += "C"
		case 'C':
			result += "G"
		case 'T':
			result += "A"
		case 'A':
			result += "U"
		}
	}

	return result
}
