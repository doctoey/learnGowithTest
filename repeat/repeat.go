package repeat

import "strings"

const repeatCount = 5

func Repeat(char string) string {
	var repeated strings.Builder
	// repeated.Grow(len(char) * repeatCount) // จองหน่วยความจำล่วงหน้าทีเดียว
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
