package pattern1_test

import (
	"github.com/murirene/go-dynamic-programming/pattern1"
	"testing"
)

func TestExercise1_1(t *testing.T) {
	inputs := []int{1, 2, 3, 4}

	if pattern1.HasSumPair(inputs) == false {
		t.Fatal("Exercise 1-1 fails")
	}
}

func TestExercise1_2(t *testing.T) {
	inputs := []int{1, 1, 3, 4, 7}

	if pattern1.HasSumPair(inputs) == false {
		t.Fatal("Exercise 1-2 fails")
	}
}

func TestExercise1_3(t *testing.T) {
	inputs := []int{2, 3, 4, 6}

	if pattern1.HasSumPair(inputs) == true {
		t.Fatal("Exercise 1-3 fails")
	}
}

func TestExercise2_1(t *testing.T) {
	inputs := []int{1, 2, 3, 7}
	target := 6

	if pattern1.HasSubsetSum(target, inputs) == false {
		t.Fatal("Exercise 2-1 fails")
	}
}

func TestExercise2_2(t *testing.T) {
	inputs := []int{1, 2, 7, 1, 5}
	target := 10

	if pattern1.HasSubsetSum(target, inputs) == false {
		t.Fatal("Exercise 2-2 fails")
	}
}

func TestExercise2_3(t *testing.T) {
	inputs := []int{1, 3, 4, 8}
	target := 6

	if pattern1.HasSubsetSum(target, inputs) == true {
		t.Fatal("Exercise 2-3 fails")
	}
}

func TestExercise3_1(t *testing.T) {
	inputs := []int{1, 2, 3, 9}
	distance := 3

	if pattern1.HasSumPairMinimumDistance(distance, inputs) == false {
		t.Fatal("Exercise 3-1 fails")
	}
}

func TestExercise3_2(t *testing.T) {
	inputs := []int{1, 2, 7, 1, 5}
	distance := 0

	if pattern1.HasSumPairMinimumDistance(distance, inputs) == false {
		t.Fatal("Exercise 3-2 fails")
	}
}

func TestExercise3_3(t *testing.T) {
	inputs := []int{1, 3, 100, 4}
	distance := 92

	if pattern1.HasSumPairMinimumDistance(distance, inputs) == true {
		t.Fatal("Exercise 3-3 fails")
	}
}

func TestExercise4_1(t *testing.T) {
	inputs := []int{1, 1, 2, 3}
	target := 4

	if pattern1.SubsetCount(target, inputs) != 3 {
		t.Fatal("Exercise 4-1 fails")
	}
}

func TestExercise4_2(t *testing.T) {
	inputs := []int{1, 2, 7, 1, 5}
	target := 9

	if pattern1.SubsetCount(target, inputs) != 3 {
		t.Fatal("Exercise 4-2 fails")
	}
}

func TestExercise5_1(t *testing.T) {
	inputs := []int{1, 1, 2, 3}
	target := 1

	if pattern1.TargetSum(target, inputs) != 3 {
		t.Fatal("Exercise 5-1 fails")
	}
}

func TestExercise5_2(t *testing.T) {
	inputs := []int{1, 2, 7, 1}
	target := 9

	if pattern1.TargetSum(target, inputs) != 2 {
		t.Fatal("Exercise 5-2 fails")
	}
}
