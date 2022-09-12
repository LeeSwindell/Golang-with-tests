package propertytests

import "strings"

type RomanNumeral struct {
	Value uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) uint16 {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(arabicNum uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabicNum >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabicNum -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var result uint16

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		numToAdd := string([]byte{symbol})
		if i < len(roman)-1 { //check we aren't at the end before making a window
			symbolDouble := roman[i:i+2] //make a window to see if its a subtractive combo
			for _, numerals := range allRomanNumerals {
				if numerals.Symbol == symbolDouble { //check the window against possible values
					numToAdd = symbolDouble
					i++
				}
			}
		}

		result += allRomanNumerals.ValueOf(numToAdd)
	}

	return result
}