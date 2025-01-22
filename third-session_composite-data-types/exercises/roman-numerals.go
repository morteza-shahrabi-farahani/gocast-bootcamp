package exercises

import (
	"fmt"
)

var romanSymbols = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input >= 4000 || input <= 0 {
		return "", fmt.Errorf("index out of range")
	}

	result := ""
	for _, symbol := range romanSymbols {
		for input >= symbol.Value {
			result += symbol.Symbol
			input -= symbol.Value
		}
	}

	return result, nil
}
