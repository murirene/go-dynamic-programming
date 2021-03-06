package pattern4_test

import (
	"github.com/murirene/go-dynamic-programming/pattern4"
	"testing"
)

func TestLPSSubsequence(t *testing.T) {
	word := "abdbca"
	if pattern4.LpsSubsequence(word) != 5 {
		t.Fatal("Failed Longest Subsequence")
	}
}

func TestLPSSubsequence2(t *testing.T) {
	word := "cddpd"
	if pattern4.LpsSubsequence(word) != 3 {
		t.Fatal("Failed Longest Subsequence")
	}
}

func TestLPSSubsequence3(t *testing.T) {
	word := "pqr"
	if pattern4.LpsSubsequence(word) != 1 {
		t.Fatal("Failed Longest Subsequence")
	}
}

func TestLPSubstring(t *testing.T) {
	word := "abdbca"
	if pattern4.LpsSubstringRecursive(word) != 3 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPSubstring2(t *testing.T) {
	word := "cddpd"
	if pattern4.LpsSubstringRecursive(word) != 3 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPSubstring3(t *testing.T) {
	word := "pqr"
	if pattern4.LpsSubstringRecursive(word) != 1 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPSubstringT1(t *testing.T) {
	word := "abdbca"
	if pattern4.LpsSubstringTabular(word) != 3 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPSubstringT2(t *testing.T) {
	word := "cddpd"
	if pattern4.LpsSubstringTabular(word) != 3 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPSubstringT3(t *testing.T) {
	word := "pqr"
	if pattern4.LpsSubstringTabular(word) != 1 {
		t.Fatal("Failed Longest Substring")
	}
}

func TestLPCount(t *testing.T) {
	word := "abdbca"
	if pattern4.PalindromeCount(word) != 7 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestLPCount2(t *testing.T) {
	word := "cddpd"
	if pattern4.PalindromeCount(word) != 7 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestLPCountT3(t *testing.T) {
	word := "pqr"
	if pattern4.PalindromeCount(word) != 3 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestMinDeletionsPalindrome(t *testing.T) {
	word := "abdbca"
	if pattern4.MinDeletionsPalindrome(word) != 1 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestMinDeletionsPalindrome2(t *testing.T) {
	word := "cddpd"
	if pattern4.MinDeletionsPalindrome(word) != 2 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestMinDeletionsPalindrome3(t *testing.T) {
	word := "pqr"
	if pattern4.MinDeletionsPalindrome(word) != 2 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestPalindromeParts(t *testing.T) {
	word := "abdbca"
	if pattern4.PalindromeParts(word) != 3 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestPalindromeParts2(t *testing.T) {
	word := "cddpd"
	if pattern4.PalindromeParts(word) != 2 {
		t.Fatal("Failed Longest Palindrome")
	}
}

func TestPalindromeParts3(t *testing.T) {
	word := "pqr"
	if pattern4.PalindromeParts(word) != 2 {
		t.Fatal("Failed Longest Palindrome")
	}
}
func TestPalindromeParts4(t *testing.T) {
	word := "pp"
	if pattern4.PalindromeParts(word) != 0 {
		t.Fatal("Failed Longest Palindrome")
	}
}


func TestPalindromeParts5(t *testing.T) {
	word := "madam"
	if pattern4.PalindromeParts(word) != 0 {
		t.Fatal("Failed Longest Palindrome")
	}
}
