/*
Write a function that takes in a Binary Tree and inverts it. In other words, the function should swap every left node in the tree for its corresponding right node.

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
       1
    /     \
   3       2
 /   \   /   \
7     6 5     4
            /   \
           9     8

*/

package trees

func (tree *BinaryTree) InvertBinaryTree() {
	if tree == nil {
		return
	}

	tree.Left.InvertBinaryTree()
	tree.Right.InvertBinaryTree()
	tree.Left, tree.Right = tree.Right, tree.Left
}
