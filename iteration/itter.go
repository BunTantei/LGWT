// Package iteration provides functions for iterating multiple time
package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

//	func Repeat(character string) string {
//		var repeated string
//		for i := 0; i < 5; i++ {
//			repeated = repeated + character
//		}
//		return repeated
//	}
//
// This is also the C way if we don't want to use the for range 5.
// ------------------------------------------
// This is an updated C way of doing it -
// const repeatCount = 5
// func Repeat(character string) string {
// var repeated string
// for i := 0; i < repeatCount; i++ {
// repeated += character
// }
// return repeatCount
// }
// ------------------------------------------
//func Repeat(character string) string {
//	var repeated string
//	for range 5 {
//		repeated = repeated + character
//	}
//	return repeated
//}
// This is the case of using for range, it's the same as the C way, just shorter and more modern.
