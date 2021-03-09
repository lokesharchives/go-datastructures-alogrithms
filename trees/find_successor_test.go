/*
Write a function that takes in a Binary Tree (where nodes have an additional pointer to their parent node) as well as a node contained in that tree and returns the given node's successor.

A node's successor is the next node to be visited (immediately after the given node) when traversing its tree using the in-order tree-traversal technique. A node has no successor if it's the last node to be visited in the in-order traversal.

If a node has no successor, your function should return None / null.

Each BinaryTree node has an integer value, a parent node, a left child node, and a right child node. Children nodes can either be BinaryTree nodes themselves or None / null.

Sample Input
tree =
              1
            /   \
           2     3
         /   \
        4     5
       /
      6
node = 5
Sample Output
1
// This tree's in-order traversal order is:
// 6 -> 4 -> 2 -> 5 -> 1 -> 3
// 1 comes immediately after 5.

*/
package trees

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func constructInorder(tree *BinaryTree) []*BinaryTree {
	if tree == nil {
		return []*BinaryTree{}
	}
	lft := constructInorder(tree.Left)
	rgt := constructInorder(tree.Right)
	return append(append(lft, tree), rgt...)
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	inOrdered := constructInorder(tree)
	for indx, data := range inOrdered {
		if data.Value == node.Value && indx+1 < len(inOrdered) {
			return inOrdered[indx+1]
		}
	}
	return nil
}

func TestFindSuccessor(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	// root.Left.Parent = root
	root.Right = &BinaryTree{Value: 3}
	// root.Right.Parent = root
	root.Left.Left = &BinaryTree{Value: 4}
	// root.Left.Left.Parent = root.Left
	root.Left.Right = &BinaryTree{Value: 5}
	// root.Left.Right.Parent = root.Left
	root.Left.Left.Left = &BinaryTree{Value: 6}
	// root.Left.Left.Left.Parent = root.Left.Left
	node := root.Left.Right
	expected := root
	actual := FindSuccessor(root, node)
	require.Equal(t, expected, actual)
}
