package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) (result string) {
	for _, v := range word {
		result = string(v) + result
	}
	return
}

// func Reverse(word string) string {
// 	var res string
// 	for _, r := range word {
// 		res = string(r) + res
// 	}
// 	return res
// }

// func Reverse(word string) string {
// 	var res string
// 	for i := len(word) - 1; i >= 0; i-- {
// 		res = res + string(word[i])
// 	}
// 	return res
// }

// func Reverse(word string) string {
// 	var res string
// 	for i := 0; i < len(word); i++ {
// 		res = string(word[i]) + res
// 	}
// 	return res
// }
