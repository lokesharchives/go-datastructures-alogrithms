package tress

import (
	"fmt"
	"math/rand"
)

// BinaryTree node
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func init() {}

// CreateRandBT Creates a Random Binary Tree
func CreateRandBT() *BinaryTree {
	var t *BinaryTree
	rand.Seed(10)

	n := 10
	for i := 0; i < n*2; i++ {
		insert(t, rand.Intn(n*2))
	}

	return t

}

// insert Inserts a node in binary tree
func insert(t *BinaryTree, v int) {
	// creates a new node with given value if the given node is nil
	if t == nil {
		t = &BinaryTree{v, nil, nil}
		return
	}

	// if the value is already in the tree then the value will not be inserted
	if v == t.Value {
		return
	}

	// nodes are inserted such that value less than current node is inserted left and
	// value greater than current node is inserted right
	if t.Value > v {
		insert(t.Left, v)
		return
	}
	insert(t.Right, v)
}

// TraverseInOrder : traverses the given tree inorder i.e., LeftNode CurrentNode RightNode
func TraverseInOrder(t *BinaryTree) {
	if t == nil {
		return
	}
	TraverseInOrder(t.Left)
	fmt.Println(t.Value)
	TraverseInOrder(t.Right)
}

// TraversePreOrder : traverses the given tree preorder i.e., CurrentNode LeftNode RightNode
func TraversePreOrder(t *BinaryTree) {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	TraversePreOrder(t.Left)
	TraversePreOrder(t.Right)
}

// TraversePostOrder : traverses the given tree postorder i.e., LeftNode RightNode CurrentNode
func TraversePostOrder(t *BinaryTree) {
	if t == nil {
		return
	}
	TraversePostOrder(t.Left)
	TraversePostOrder(t.Right)
	fmt.Println(t.Value)
}
