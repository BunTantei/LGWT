// Package iteration provides functions for iterating multiple time
package iteration

func Repeat(character string) string {
	var repeated string
	for range 5 {
		repeated = repeated + character
	}
	return repeated
}

// func Repeat(character string) string {
// 	var repeated string
// 	for i := 0; i < 5; i++ {
// 		repeated = repeated + character
// 	}
//	return repeated
// }
// This is also the C way if we don't want to use the for range 5.
