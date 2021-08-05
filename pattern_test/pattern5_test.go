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
