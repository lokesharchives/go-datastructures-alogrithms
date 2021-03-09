/*
Write a function that takes in a non-empty array of integers that are sorted in ascending order and returns a new array of the same length with the squares of the original integers also sorted in ascending order.

Sample Input
array = [1, 2, 3, 5, 6, 8, 9]
Sample Output
[1, 4, 9, 25, 36, 64, 81]
*/

package arrays

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// sortedSquaredArray return square of elements of given sorted array
// o(nlogn) time and o(1) space
func SortedSquaredArray(arr []int) []int {
	// Write your code here.
	// replace all the negative number with positive numbers
	indx := 0
	for indx < len(arr) && arr[indx] < 0 {
		arr[indx] = -arr[indx]
		indx++
	}

	// sort the given array o(nlogn)
	sort.Ints(arr)

	// square the elements of the array
	for indx, data := range arr {
		arr[indx] = data * data
	}

	return arr
}

// TestSortedSquaredArray test case for sorted squated array
func TestSortedSquaredArray(t *testing.T) {

	arr := []int{-4, -2, 0, 1, 2, 3, 6, 10}
	require.Equal(t, []int{0, 1, 4, 4, 9, 16, 36, 100}, SortedSquaredArray(arr))
}
