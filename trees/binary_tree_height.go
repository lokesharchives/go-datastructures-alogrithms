/*
Height Balanced Binary Tree
You're given the root node of a Binary Tree. Write a function that returns true if this Binary Tree is height balanced and false if it isn't.

A Binary Tree is height balanced if for each node in the tree, the difference between the height of its left subtree and the height of its right subtree is at most 1.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree = 1
     /   \
    2     3
  /   \     \
 4     5     6
     /   \
    7     8
*/
package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func heightOfTree(tree *BinaryTree) (int, bool) {
	if tree == nil {
		return 0, true
	}
	lftHight, lft := heightOfTree(tree.Left)
	rgtHight, rgt := heightOfTree(tree.Right)
	return 1 + max(lftHight, rgtHight), lft && rgt && abs(lftHight-rgtHight) <= 1
}

// HeightBalancedBinaryTree logic
func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	// Write your code here.
	_, res := heightOfTree(tree)
	return res
}

// TestHeightBalancedBinaryTree tests
func TestHeightBalancedBinaryTree(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Left.Right.Left = &BinaryTree{Value: 7}
	root.Left.Right.Right = &BinaryTree{Value: 8}
	expected := true
	actual := HeightBalancedBinaryTree(root)
	require.Equal(t, expected, actual)
}
