package main

import (
	"strings"
)

// Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
func EvenOrOdd(number int) string {
	if number % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}



// Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.
func BoolToWord(word bool) string {
	if word {
		return "Yes"
	} else {
		return "No"
	}
}

//Create a function that removes the first and last characters of a string.
//You're given one parameter, the original string.
//You don't have to worry with strings with less than two characters.
func RemoveChar(word string) string {
	return string(word[1:len(word)-1])
}


// Complete the square sum function so that it squares each number passed into it and then sums the results together.
func SquareSum(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v*v
	}
	return result
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func Move(position, roll int) int {
	return position + roll*2
}

// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
func AbbrevName(name string) string {
	initials := strings.SplitAfter(name, " ")                                     //make a slice
	return strings.ToUpper(string(initials[0][0]) + "." + string(initials[1][0])) // <- use this syntax to access the value of elements in a slice by index
}

// Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
func Invert(arr []int) []int {
	var result []int
	for _, v := range arr {
		result = append(result, -v)
	}
	return result
}
