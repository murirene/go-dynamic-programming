package pattern3_test

import (
	"github.com/murirene/go-dynamic-programming/pattern3"
	"testing"
)

func TestFib(t *testing.T) {
	if pattern3.Fib(5) != 5 {
		t.Fatal("Failed 3-1")
	}

	if pattern3.Fib(6) != 8 {
		t.Fatal("Failed 3-1")
	}

	if pattern3.Fib(7) != 13 {
		t.Fatal("Failed 3-1")
	}
}

func TestStaircase(t *testing.T) {
	if pattern3.StaircaseRecursive(4) != 7 {
		t.Fatal("Fatal 3-2")
	}

	if pattern3.StaircaseRecursive(5) != 13 {
		t.Fatal("Fatal 3-3")
	}
}

func TestStaircaseTabular(t *testing.T) {
	if pattern3.StaircaseTabular(4) != 7 {
		t.Fatal("Fatal 3-2 tabular")
	}

	if pattern3.StaircaseTabular(5) != 13 {
		t.Fatal("Fatal 3-3 tabular")
	}
}
