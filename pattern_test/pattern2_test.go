package pattern2_test

import (
	"github.com/murirene/go-dynamic-programming/pattern2"
	"testing"
)

func TestExercise1_1(t *testing.T) {
	profits := []int{15, 50, 60, 90}
	weights := []int{1, 3, 4, 5}
	capacity := 8

	if pattern2.UnboundKnapsack(weights, profits, capacity) != 140 {
		t.Fatal("Exercise 1-1 failed")
	}

}

func TestExercise1_2(t *testing.T) {
	profits := []int{15, 20, 50}
	weights := []int{1, 2, 3}
	capacity := 5

	if pattern2.UnboundKnapsack(weights, profits, capacity) != 80 {
		t.Fatal("Exercise 1-2 failed")
	}

}
