package main

import (
	"strings"
)

// Complete the square sum function so that it squares each number passed into it and then sums the results together.
func SquareSum(numbers []int) int {
    // your code here
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
