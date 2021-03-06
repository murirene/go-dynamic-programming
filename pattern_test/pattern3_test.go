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

func TestStairCaseReachEndTab(t *testing.T) {
	list := []int{2, 1, 1, 1, 4}
	if pattern3.MinReachEndTabular(list) != 3 {
		t.Fatal("Fatal 5-1 Tab")
	}
}

func TestStairCaseReachEnd2Tab(t *testing.T) {
	list := []int{1, 1, 3, 6, 9, 3, 0, 1, 3}
	if pattern3.MinReachEndTabular(list) != 4 {
		t.Fatal("Fatal 5-2 Tab")
	}
}

func TestStairCaseMinFeeTabularT1(t *testing.T) {
	steps := []int{1, 2, 5, 2, 1, 2}
	if pattern3.MinFeeReachEnd(steps) != 3 {
		t.Fatal("Failed Case 1")
	}
}

func TestStairCaseMinFeeTabularT2(t *testing.T) {
	steps := []int{2, 3, 4, 5}
	if pattern3.MinFeeReachEnd(steps) != 5 {
		t.Fatal("Failed Case 2")
	}
}

func TestHouseTheif(t *testing.T) {
    houses := []int{2, 5, 1, 3, 6, 2, 4}
    if pattern3.HouseTheif(houses) != 15 {
        t.Fatal("Failed House thief 1")
    }
}


func TestHouseTheif2(t *testing.T) {
    houses := []int{2, 10, 14, 8, 1}
    if pattern3.HouseTheif(houses) != 18 {
        t.Fatal("Failed House thief 2")
    }
}
