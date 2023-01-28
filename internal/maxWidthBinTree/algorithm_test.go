package maxwidthbintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWidthOfBinaryTree(t *testing.T) {
	binTree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := 2

	result := widthOfBinaryTree(binTree)

	assert.Equal(t, expected, result, "difference on example!")
}
