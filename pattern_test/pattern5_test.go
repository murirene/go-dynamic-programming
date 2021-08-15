package pattern_test

import (
	"fmt"
	"github.com/murirene/go-dynamic-programming/pattern5"
	"testing"
)

func TestLCS1(t *testing.T) {
	word1 := "abdca"
	word2 := "cbda"
	result := pattern5.LcSubstring(word1, word2)
	if result != 2 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCS2(t *testing.T) {
	word1 := "passport"
	word2 := "ppsspt"
	result := pattern5.LcSubstring(word1, word2)

	if result != 3 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCST1(t *testing.T) {
	word1 := "abdca"
	word2 := "cbda"
	result := pattern5.LcSubstringTabular(word1, word2)
	if result != 2 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCST2(t *testing.T) {
	word1 := "passport"
	word2 := "ppsspt"
	result := pattern5.LcSubstringTabular(word1, word2)

	if result != 3 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSsT1(t *testing.T) {
	word1 := "abdca"
	word2 := "cbda"
	result := pattern5.LcSubsequence(word1, word2)
	if result != 3 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSs2(t *testing.T) {
	word1 := "passport"
	word2 := "ppsspt"
	result := pattern5.LcSubsequence(word1, word2)

	if result != 5 {
		t.Fatal("Failed Longest Common Substring " + fmt.Sprintf("%d\n", result))
	}
}

func TestMinEdits1(t *testing.T) {
	word1 := "abc"
	word2 := "fbc"
	deletes := 1
	inserts := 1

	deletes, inserts = pattern5.LcMinimalEdits(word1, word2)

	if deletes != 1 && inserts != 1 {
		t.Fatal("Failed Minimal Edits " + fmt.Sprintf("edits %d inserts %d\n", deletes, inserts))
	}
}

func TestMinEdits2(t *testing.T) {
	word1 := "abdca"
	word2 := "cbda"

	deletes, inserts := pattern5.LcMinimalEdits(word1, word2)

	if deletes != 2 && inserts != 1 {
		t.Fatal("Failed Minimal Edits " + fmt.Sprintf("edits %d inserts %d\n", deletes, inserts))
	}
}

func TestMinEdits3(t *testing.T) {
	word1 := "passport"
	word2 := "ppsspt"

	deletes, inserts := pattern5.LcMinimalEdits(word1, word2)

	if deletes != 3 && inserts != 1 {
		t.Fatal("Failed Minimal Edits " + fmt.Sprintf("edits %d inserts %d\n", deletes, inserts))
	}
}

func TestLongestIncreasingSubsequence4(t *testing.T) {
	list := []int{4, 2, 3, 6, 10, 1, 12}
	result := pattern5.Lis(list)

	if result != 5 {
		t.Fatal("Failed Longest Increasing Subsequence  " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestIncreasingSubsequence(t *testing.T) {
	list := []int{-4, 10, 3, 7, 15}
	result := pattern5.Lis(list)

	if result != 4 {
		t.Fatal("Failed Longest Increasing Subsequence  " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestIncreasingSubsequence4Tab(t *testing.T) {
	list := []int{4, 2, 3, 6, 10, 1, 12}
	result := pattern5.LisTabular(list)

	if result != 5 {
		t.Fatal("Failed Longest Increasing Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestIncreasingSubsequenceTab2(t *testing.T) {
	list := []int{-4, 10, 3, 7, 15}
	result := pattern5.LisTabular(list)

	if result != 4 {
		t.Fatal("Failed Longest Increasing Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestIncreasingSubsequenceRedux(t *testing.T) {
	list := []int{4, 2, 3, 6, 10, 1, 12}
	result := pattern5.LongestIncreasingSequence(list)

	if result != 5 {
		t.Fatal("Failed Longest Increasing Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestIncreasingSubsequenceRedux2(t *testing.T) {
	list := []int{-4, 10, 3, 7, 15}
	result := pattern5.LongestIncreasingSequence(list)

	if result != 4 {
		t.Fatal("Failed Longest Increasing Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestSumSequence(t *testing.T) {
	list := []int{4, 1, 3, 6, 10, 1, 12}
	result := pattern5.LongestSumSequence(list)

	if result != 32 {
		t.Fatal("Failed Longest Sum Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestSumSubsequence2(t *testing.T) {
	list := []int{-4, 10, 3, 7, 15}
	result := pattern5.LongestSumSequence(list)

	if result != 25 {
		t.Fatal("Failed Longest Sum Subsequence Tab " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestCommonSuperSequence(t *testing.T) {
	word1 := "abcf"
	word2 := "bdcf"
	result := pattern5.ShortestCommonSuperSequence(word1, word2)

	if result != 5 {
		t.Fatal("Failed Longest Common Super Subsequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLongestCommonSuperSubsequence2(t *testing.T) {
	word1 := "dynamic"
	word2 := "programming"
	result := pattern5.ShortestCommonSuperSequence(word1, word2)

	if result != 15 {
		t.Fatal("Failed Longest Common Super Subsequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice(t *testing.T) {
	word1 := "a"
	word2 := "a"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 1 {
		t.Fatal("Failed 1 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice2(t *testing.T) {
	word1 := "ab"
	word2 := "fb"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 3 {
		t.Fatal("Failed 2 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice3(t *testing.T) {
	word1 := "ba"
	word2 := "ab"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 3 {
		t.Fatal("Failed 3 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice4(t *testing.T) {
	word1 := "abcf"
	word2 := "bdcf"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 5 {
		t.Fatal("Failed 3 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice5(t *testing.T) {
	word1 := "a"
	word2 := ""
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 1 {
		t.Fatal("Failed 3 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice6(t *testing.T) {
	word1 := ""
	word2 := "a"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 1 {
		t.Fatal("Failed 3 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestLCSRecursice7(t *testing.T) {
	word1 := "dynamic"
	word2 := "programming"
	result := pattern5.ShortestSuperSequenceRecursive(word1, word2, 0, 0)

	if result != 15 {
		t.Fatal("Failed 3 to find shortest super sequence " + fmt.Sprintf("%d\n", result))
	}
}

func TestMinimumDeletes(t *testing.T) {
	list := []int{4, 2, 3, 6, 10, 1, 12}
	result := pattern5.MinimumDeletes(list, -1, 0)
	if result != 2 {
		t.Fatal("Failed problem 1 " + fmt.Sprintf("%d\n", result))
	}
}

func TestMinimumDeletes2(t *testing.T) {
	list := []int{-4, 10, 3, 7, 15}
	result := pattern5.MinimumDeletes(list, -1, 0)
	if result != 1 {
		t.Fatal("Failed problem 1 " + fmt.Sprintf("%d\n", result))
	}
}

func TestMinimumDeletes3(t *testing.T) {
	list := []int{3, 2, 1, 0}
	result := pattern5.MinimumDeletes(list, -1, 0)
	if result != 3 {
		t.Fatal("Failed problem 1 " + fmt.Sprintf("%d\n", result))
	}
}

func TestShortestCommonSuperSequence1(t *testing.T) {
	word1 := "abcf"
	word2 := "bdcf"
	result := pattern5.ShortestCommonSuperSequence(word1, word2)
	if result != 5 {
		t.Fatal("Failed ShortestCommonSuperSequence " + fmt.Sprintf("%s %s = %d\n", word1, word2, result))
	}
}

func TestShortestCommonSuperSequence2(t *testing.T) {
	word1 := "dynamic"
	word2 := "programming"
	result := pattern5.ShortestCommonSuperSequence(word1, word2)
	if result != 15 {
		t.Fatal("Failed ShortestCommonSuperSequence " + fmt.Sprintf("%s %s = %d\n", word1, word2, result))
	}
}
