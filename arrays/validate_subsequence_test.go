/*
Given two non-empty arrays of integers, write a function that determines whether the second array is a subsequence of the first one.

A subsequence of an array is a set of numbers that aren't necessarily adjacent in the array but that are in the same order as they appear in the array. For instance, the numbers [1, 3, 4] form a subsequence of the array [1, 2, 3, 4], and so do the numbers [2, 4]. Note that a single number in an array and the array itself are both valid subsequences of the array.

Sample Input
array = [5, 1, 22, 25, 6, -1, 8, 10]
sequence = [1, 6, -1, 10]
Sample Output
true
*/
package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// validateSubsequence o(n) time and o(1) space
func validateSubsequence(arr, seq []int) bool {
	arrIndx, seqIndx := 0, 0
	for arrIndx < len(arr) && seqIndx < len(seq) {
		if arr[arrIndx] == seq[seqIndx] {
			seqIndx++
		}
		arrIndx++
	}

	return seqIndx == len(seq)
}

func TestValidateSubsequence(t *testing.T) {
	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, 10}
	require.Equal(t, true, validateSubsequence(arr, seq))
}
