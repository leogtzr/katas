package main

import "strings"

var (
	units = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen",
	}
	tens = []string{
		"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}
)

func numToWord(num int) string {
	if num == 0 {
		return units[0]
	} else if num < 20 {
		return units[num]
	} else if num < 100 {
		var numWord strings.Builder

		numWord.WriteString(tens[num/10])

		if num%10 != 0 {
			numWord.WriteString(" ")
			numWord.WriteString(units[num%10])
		}

		return numWord.String()
	} else if num < 1000 {
		var sb strings.Builder

		sb.WriteString(units[num/100])
		sb.WriteString(" hundred")

		if num%100 != 0 {
			sb.WriteString(" and ")
			sb.WriteString(numToWord(num % 100))
		}

		return sb.String()
	} else if num < 1000000 {
		var numWord strings.Builder

		numWord.WriteString(numToWord(num / 1000))
		numWord.WriteString(" thousand")

		if num%1000 != 0 {
			if num%1000 < 100 {
				numWord.WriteString(" and ")
			} else {
				numWord.WriteString(" ")
			}
			numWord.WriteString(numToWord(num % 1000))
		}

		return numWord.String()
	} else {
		return "Number out of range"
	}
}
