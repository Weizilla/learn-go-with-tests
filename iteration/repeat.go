package iteration

import "strings"

func Repeat(value string, num int) string {
	var repeated strings.Builder
	for i := 0; i < num; i++ {
		repeated.WriteString(value)
	}
	return repeated.String()
}
