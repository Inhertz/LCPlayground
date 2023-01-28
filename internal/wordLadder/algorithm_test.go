package wordladder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	expected := 5

	result := ladderLength(beginWord, endWord, wordList)

	assert.Equal(t, expected, result, "error on test ladder length")
}
