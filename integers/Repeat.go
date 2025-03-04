package integers

import "strings"

const repeatCount = 5

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += repeated + character // this is not efficient, especially as mem for it grows
// 	}
// 	return repeated
// }

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character) // this is not efficient, especially as mem for it grows
	}
	return repeated.String()
}

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a")
// 	}
// }
