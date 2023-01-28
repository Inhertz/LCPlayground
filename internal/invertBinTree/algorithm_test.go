package invertbintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {

	leftNode := &TreeNode{Val: 2}
	rightNode := &TreeNode{Val: 3}
	tree := &TreeNode{Val: 1, Left: leftNode, Right: rightNode}
	expected := &TreeNode{Val: 1, Left: rightNode, Right: leftNode}

	result := invertTree(tree)

	assert.Equal(t, expected, result, "values should be equal")
}
