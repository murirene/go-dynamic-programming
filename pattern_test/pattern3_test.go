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

func TestCountWays(t *testing.T) {
	if pattern3.CountWays(6) != 9 {
		t.Fatal("Fatal 4-1")
	}
}

func TestCountWaysTabular(t *testing.T) {
	if pattern3.CountWaysTabular(5) != 6 {
		t.Fatal("Fatal 4-2")
	}

	if pattern3.CountWaysTabular(6) != 9 {
		t.Fatal("Fatal 4-3")
	}
}

func TestStairCaseReachEnd(t *testing.T) {
	list := []int{2, 1, 1, 1, 4}
	if pattern3.MinReachEnd(list) != 3 {
		t.Fatal("Fatal 5-1")
	}
}

func TestStairCaseReachEnd2(t *testing.T) {
	list := []int{1, 1, 3, 6, 9, 3, 0, 1, 3}
	if pattern3.MinReachEnd(list) != 4 {
		t.Fatal("Fatal 5-1")
	}
}
