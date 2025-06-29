// Package iteration provides functions for iterating multiple time
package iteration

func Repeat(character string) string {
	var repeated string
	for range 5 {
		repeated = repeated + character
	}
	return repeated
}
