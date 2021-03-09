/*
Write a function that takes in a Binary Tree and returns its max path sum.

A path is a collection of connected nodes in a tree, where no node is connected to more than two other nodes; a path sum is the sum of the values of the nodes in a particular path.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree = 1
    /     \
   2       3
 /   \   /   \
4     5 6     7
Sample Output
18 // 5 + 2 + 1 + 3 + 7

*/
package trees

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type MaxPath struct {
	maxPathSum, pathSum int
}

func getMaxpath(tree *BinaryTree) MaxPath {
	if tree == nil {
		return MaxPath{math.MinInt32, 0}
	}

	lft := getMaxpath(tree.Left)
	rgt := getMaxpath(tree.Right)

	maxTillNow := max(max(lft.maxPathSum, rgt.maxPathSum), lft.pathSum+rgt.pathSum+tree.Value)
	maxPath := max(lft.pathSum, rgt.pathSum) + tree.Value

	return MaxPath{max(maxTillNow, maxPath), maxPath}
}

func MaxPathSum(tree *BinaryTree) int {
	maxPath := getMaxpath(tree)
	return maxPath.maxPathSum
}

func TestMaxPathSum(t *testing.T) {
	test := NewBinaryTree(2, 3, 4, 5, 6, 7)
	require.Equal(t, MaxPathSum(test), 22)
}
