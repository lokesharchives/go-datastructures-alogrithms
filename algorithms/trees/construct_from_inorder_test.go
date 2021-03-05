package trees

import (
	"testing"

	trees "github.com/lokesharchives/go-datastructures-alogrithms/datastructures/trees"
	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	assert.True(t, true, "True is true!")

	var bt *trees.BinaryTree
	bt = trees.CreateRandBT()
	trees.TraverseInOrder(bt)
}
