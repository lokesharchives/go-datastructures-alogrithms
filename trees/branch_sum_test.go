/*
Write a function that takes in a Binary Tree and returns a list of its branch sums ordered from leftmost branch sum to rightmost branch sum.

A branch sum is the sum of all values in a Binary Tree branch. A Binary Tree branch is a path of nodes in a tree that starts at the root node and ends at any leaf node.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree =     1
        /     \
       2       3
     /   \    /  \
    4     5  6    7
  /   \  /
 8    9 10
Sample Output
[15, 16, 18, 10, 11]
// 15 == 1 + 2 + 4 + 8
// 16 == 1 + 2 + 4 + 9
// 18 == 1 + 2 + 5 + 10
// 10 == 1 + 3 + 6
// 11 == 1 + 3 + 7
*/

package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BranchSums(root *BinaryTree) []int {
	if root == nil {
		return []int{}
	}

	res := append(BranchSums(root.Left), BranchSums(root.Right)...)
	if len(res) == 0 {
		res = append(res, root.Value)
		return res
	}
	for i := 0; i < len(res); i++ {
		res[i] += root.Value
	}

	return res
}

//TestBranchSum test case 1
func TestBranchSum(t *testing.T) {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	expected := []int{15, 16, 18, 10, 11}
	output := BranchSums(tree)
	require.Equal(t, expected, output)
}
