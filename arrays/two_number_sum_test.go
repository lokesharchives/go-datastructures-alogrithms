/*
Write a function that takes in a non-empty array of distinct integers and an integer representing a target sum. If any two numbers in the input array sum up to the target sum, the function should return them in an array, in any order. If no two numbers sum up to the target sum, the function should return an empty array.

Note that the target sum has to be obtained by summing two different integers in the array; you can't add a single integer to itself in order to obtain the target sum.

You can assume that there will be at most one pair of numbers summing up to the target sum.

Sample Input
array = [3, 5, -4, 8, 11, 1, -1, 6]
targetSum = 10
Sample Output
[-1, 11] // the numbers could be in reverse order
*/
package arrays

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// twoNumberSumByHashMap 0(n) space and 0(1) time
func twoNumberSumByHashMap(arr []int, target int) []int {
	hashmap := make(map[int]bool)

	for _, data := range arr {
		diff := target - data
		if _, ok := hashmap[diff]; ok {
			return []int{data, diff}
		}
		hashmap[data] = true
	}

	return []int{}

}

// twoNumberSumBySorting find two number sum with o(nlogn) time and o(1) space
func twoNumberSumBySorting(arr []int, target int) []int {

	//Sorting the arry o(nlogn)
	sort.Ints(arr)

	// Pointer lft pointing to the start of the array
	// Pointer rgt pointing to the end of the arry
	lft := 0
	rgt := len(arr) - 1

	// Loop till lft less than right
	for lft < rgt {
		// if the sum of lft and rgt pointer arr values is equal to
		// given sum then we return the lft and rgt array values
		sum := arr[lft] + arr[rgt]
		if sum == target {
			return []int{arr[lft], arr[rgt]}
		}

		// if the left and right pointer sum is greater than the target
		// We will move the rgt pointer to reduce the sum
		if sum > target {
			rgt--
			continue
		}

		// else we will increase the lft pointer to decrease the sum
		lft++
	}

	// When no two elements sum up to the given target return empty array
	return []int{}
}

func TestTwoNumberSumBySorting(t *testing.T) {

	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10

	require.Equal(t, []int{-1, 11}, twoNumberSumBySorting(arr, target))
}

func TestTwoNumberSumByHashMap(t *testing.T) {

	arr := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10

	require.Equal(t, []int{-1, 11}, twoNumberSumByHashMap(arr, target))
}
