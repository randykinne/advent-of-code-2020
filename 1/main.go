package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*
	Summary of answer:
	target = 2020
	We take half of 2020 (1010) and divide the inputs into greater than and lower than,
	 as any answer will have one in each that add up to target

	 then we subtract the lower values from 2020 to get values with one that matches an upper value,
	 and then search each result in subtracted lower with each upper until we have our two values that sum to target

	 Then we multiply each of those values together for the Advent of Code answer
*/
func main() {
	rawData := ReadInput()
	data := ConvertByteArrToIntArr(rawData)

	// Magic number for advent of code
	target := 2020
	upper, lower := SplitIntArrToUpperAndLower(target, data)

	subtracted := MapSubtractValuesFromTarget(target, lower)

	val1, val2 := FindMatchingIndexPair(target, upper, subtracted)

	fmt.Printf("Value 1: " + strconv.Itoa(upper[val1]) + " & Value 2: " + strconv.Itoa(lower[val2]))
	fmt.Printf("\nResult for Advent of Code: " + strconv.Itoa(upper[val1]*lower[val2]))
}

// FindMatchingIndexPair of target-i1 values that equal i2 values, returns the indices that contain matching values
func FindMatchingIndexPair(target int, i1 []int, i2 []int) (int1 int, int2 int) {
	for i, item1 := range i1 {
		for j, item2 := range i2 {
			if item1 == item2 {
				return i, j
			}
		}
	}
	return 0, 0
}

// MapSubtractValuesFromTarget returns an array that has been subtracted
// like a map function, mapping a subtract to each value and returning the arr
func MapSubtractValuesFromTarget(target int, ints []int) []int {
	result := make([]int, len(ints))
	for i, val := range ints {
		result[i] = target - val
	}
	return result
}

// SplitIntArrToUpperAndLower returns two arrays that represent the upper and lower based on target
func SplitIntArrToUpperAndLower(target int, ints []int) ([]int, []int) {
	lower := make([]int, 0)
	upper := make([]int, 0)
	for _, val := range ints {
		if val <= target/2 {
			upper = append(upper, val)
		} else {
			lower = append(lower, val)
		}
	}
	return lower, upper
}

// ConvertByteArrToIntArr converts a byte arr to int arr
func ConvertByteArrToIntArr(b []byte) []int {
	lines := strings.Split(string(b), "\n")
	nums := make([]int, len(lines))
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		Check(err)
		nums[i] = n
	}
	return nums
}

// ReadInput from input file
func ReadInput() []byte {
	data, err := ioutil.ReadFile("input")
	Check(err)
	return data
}

// Check for errors
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
