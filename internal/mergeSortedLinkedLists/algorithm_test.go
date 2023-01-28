package mergesortedlinkedlists

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}
	list2 := &ListNode{Val: 2, Next: nil}
	expected := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}

	result := mergeTwoLists(list1, list2)

	assert.Equal(t, expected, result, fmt.Sprintf("Expected : %v, Got: %v", expected, result))

}
