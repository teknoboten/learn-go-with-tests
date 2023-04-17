package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}


// func Sum(numbers [5]int) int {
// 	result := 0
	
// 	for _, number := range numbers {
// 		fmt.Println(number)
// 		result += number
// 	}
	
// 	return result
// }



// func Sum(numbers [5]int) int {
// 	result := 0
// 	for i := 0; i < 5; i++ {
// 		result += numbers[i]
// 	}
// 	return result
// }