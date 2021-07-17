package roman_numerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var all = RomanNumerals{
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range all {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToArabic(roman string) uint16 {
	var total uint16

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(roman, i, symbol) {
			if value := all.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
				continue
			}
		}

		total += all.ValueOf(symbol)
	}
	return total
}

func couldBeSubtractive(roman string, i int, symbol uint8) bool {
	isSubtractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return i+1 < len(roman) && isSubtractiveSymbol
}
