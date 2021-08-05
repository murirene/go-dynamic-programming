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
