package main

import "fmt"

func romanToDecimalInefficient(roman string) int {
	var romanNumerals = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := romanNumerals[roman[i]]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

func decimalToRoman(decimal int) string {
	var romanNumerals = []struct {
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

	roman := ""
	for _, numeral := range romanNumerals {
		for decimal >= numeral.Value {
			roman += numeral.Symbol
			decimal -= numeral.Value
		}
	}
	return roman
}

func main() {
	fmt.Println(decimalToRoman(romanToDecimalInefficient("XXXXX")))
	fmt.Println(decimalToRoman(romanToDecimalInefficient("XXXXXVVII")))
}
