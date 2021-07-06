package pattern1_test

import (
	"github.com/murirene/go-dynamic-programming/pattern1"
	"testing"
)

func TestExercise1(t *testing.T) {
	inputs := []int{1, 2, 3, 4}

	if pattern1.HasSumPair(inputs) == false {
		t.Fatal("Exercise 1 fails")
	}
}

func TestExercise2(t *testing.T) {
	inputs := []int{1, 1, 3, 4, 7}

	if pattern1.HasSumPair(inputs) == false {
		t.Fatal("Exercise 2 fails")
	}
}

func TestExercise3(t *testing.T) {
	inputs := []int{2, 3, 4, 6}

	if pattern1.HasSumPair(inputs) == true {
		t.Fatal("Exercise 3 fails")
	}
}
