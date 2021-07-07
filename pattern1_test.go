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


