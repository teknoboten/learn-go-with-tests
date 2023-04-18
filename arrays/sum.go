package main

//takes in a int slice 'numbers' and returns the sum of all values
func Sum(numbers []int) int {
	sum := 0

//iterates over numbers using a range, which returns the index and value (number) on each iteration
//use the blank identifier to ignore the index
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//takes in a slice of integer arrays (it's a variadic function!)
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum) //len returns the length
	sums := make([]int, lengthOfNumbers) //use make to initalize a slice with the appropriate length

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

//iterates over a collection of slices and returns the sum of each 'tail' (everything but the first value)
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]            //each tail is defined using a slice of a slice
			sums = append(sums, Sum(tail)) //append the resulting slice to the sums slice
		}
	}
	return sums
}
