package layout

import (
	"regexp"
)

func formatCommas(num string) string {
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for {
		formatted := re.ReplaceAllString(num, "$1,$2")
		if formatted == num {
			return formatted
		}
		num = formatted
	}
}
