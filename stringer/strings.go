package stringer

import "strings"

func Stringer(s, old, new string) string {
	replaced := strings.ReplaceAll(s, old, new)
	result := strings.ToLower(replaced)
	return result
}


// func Stringer(x, y string) bool {
// 	var result bool
// 	if strings.EqualFold(x, y) {
// 		result = true
// 	}
// 	return result
// }

// func Stringer(x, y string) bool {
// 	var result bool
// 	if strings.EqualFold(x, y) {
// 		result = true
// 	}
// 	return result
// }





// func Stringer(x, y string) string {
// 	var result string

// 	if strings.Compare(x, y) == 0 {
// 		result = "match!"
// 	}
// 	if strings.Compare(x, y) != 0 {
// 		result = "not a match!"
// 	}
// 	return result
// }