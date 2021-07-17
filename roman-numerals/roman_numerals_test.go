package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var tests = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{500, "D"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d->%q", test.arabic, test.roman), func(t *testing.T) {
			got := ConvertToRoman(test.arabic)
			want := test.roman

			if got != want {
				t.Errorf("arabic %q, roman %q", got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("%q->%d", test.roman, test.arabic), func(t *testing.T) {
			got := ConvertToArabic(test.roman)
			want := test.arabic

			if got != want {
				t.Errorf("arabic %d, roman %q", got, test.roman)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
