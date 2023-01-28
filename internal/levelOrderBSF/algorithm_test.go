package levelorderbsf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	binTree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := [][]int{{3}, {9, 20}, {15, 7}}

	result := levelOrder(binTree)

	assert.Equal(t, expected, result, "difference on example!")
}
