package avglevelsbintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageOfLevels(t *testing.T) {
	binTree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := []float64{3, 14.5, 11}

	result := averageOfLevels(binTree)

	assert.Equal(t, expected, result, "difference on example!")
}
