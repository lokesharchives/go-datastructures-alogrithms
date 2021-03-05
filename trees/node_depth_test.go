/*
The distance between a node in a Binary Tree and the tree's root is called the node's depth.

Write a function that takes in a Binary Tree and returns the sum of its nodes' depths.

Each BinaryTree node has an integer value, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree =    1
       /     \
      2       3
    /   \   /   \
   4     5 6     7
 /   \
8     9
Sample Output
16
// The depth of the node with value 2 is 1.
// The depth of the node with value 3 is 1.
// The depth of the node with value 4 is 2.
// The depth of the node with value 5 is 2.
// Etc..
// Summing all of these depths yields 16.

*/
package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func getNodeDepth(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}
	return depth + getNodeDepth(root.Left, depth+1) + getNodeDepth(root.Right, depth+1)
}

// NodeDepths returns the sum of depth of all the nodes
func NodeDepths(root *BinaryTree) int {
	return getNodeDepth(root, 0)
}

func TestNodeDepth(t *testing.T) {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9)
	require.Equal(t, 16, NodeDepths(tree))
}
