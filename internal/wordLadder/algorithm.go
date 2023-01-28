package wordladder

func ladderLength(beginWord string, endWord string, wordList []string) int {

	dic := make(map[string]int)
	for i, word := range wordList {
		dic[word] = i
	}

	if _, ok := dic[endWord]; !ok {
		return 0
	}

	return 0
}
