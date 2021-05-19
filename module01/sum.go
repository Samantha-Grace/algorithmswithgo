package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	//var result int
	result := 0
	for _, v := range numbers {
		//result = result + val
		result += v
	}
	return result
}

// Another way to do this: recursion
// func Sum(numbers []int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}
// 	return numbers[0] + Sum(numbers[1:])
// }
