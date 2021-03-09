/*
Write a function that takes in a Binary Tree and returns its diameter. The diameter of a binary tree is defined as the length of its longest path, even if that path doesn't pass through the root of the tree.

A path is a collection of connected nodes in a tree, where no node is connected to more than two other nodes. The length of a path is the number of edges between the path's first node and its last node.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree =        1
            /   \
           3     2
         /   \
        7     4
       /       \
      8         5
     /           \
    9             6
Sample Output
6 // 9 -> 8 -> 7 -> 3 -> 4 -> 5 -> 6
// There are 6 edges between the
// first node and the last node
// of this tree's longest path.
*/

package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getPathLength(tree *BinaryTree, diameter *int) int {
	if tree == nil {
		return 0
	}
	lmaxPath := getPathLength(tree.Left, diameter)
	rmaxPath := getPathLength(tree.Right, diameter)
	*diameter = max(lmaxPath+rmaxPath, *diameter)
	return max(lmaxPath, rmaxPath) + 1

}

func BinaryTreeDiameter(tree *BinaryTree) int {
	var diameter int = 0
	getPathLength(tree, &diameter)
	return diameter
}

func TestBinaryTreeDiameter(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 7}
	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Left.Left.Left = &BinaryTree{Value: 9}
	root.Left.Right = &BinaryTree{Value: 4}
	root.Left.Right.Right = &BinaryTree{Value: 5}
	root.Left.Right.Right.Right = &BinaryTree{Value: 6}
	root.Right = &BinaryTree{Value: 2}
	expected := 6
	actual := BinaryTreeDiameter(root)
	require.Equal(t, expected, actual)
}
