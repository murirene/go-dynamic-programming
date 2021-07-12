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

func TestExercise2_1(t *testing.T) {
	profits := []int{2, 6, 7, 10, 13}
	lengths := []int{1, 2, 3, 4, 5}
	capacity := 5

	if pattern2.RodCutOptimalProfit(lengths, profits, capacity) != 14 {
		t.Fatal("Exercise 2-1 failed")
	}

}

func TestExercise3_1(t *testing.T) {
	denominations := []int{1, 2, 3}
	amount := 5

	if pattern2.CoinChangeCount(denominations, amount) != 5 {
		t.Fatal("Exercise 3-1 failed")
	}

}

func TestExercise4_1(t *testing.T) {
	denominations := []int{1, 2, 3}
	amount := 5

	if pattern2.MinimumCoinChangeCount(denominations, amount) != 2 {
		t.Fatal("Exercise 4-1 failed")
	}

}

func TestExercise5_1(t *testing.T) {
	ribbonLengths := []int{2, 3, 5}
	cut := 5

	if pattern2.MaxRibbonCut(ribbonLengths, cut) != 2 {
		t.Fatal("Exercise 5-1 failed")
	}
}

func TestExercise5_2(t *testing.T) {
	ribbonLengths := []int{2, 3}
	cut := 7

	if pattern2.MaxRibbonCut(ribbonLengths, cut) != 3 {
		t.Fatal("Exercise 5-2 failed")
	}
}

func TestExercise5_3(t *testing.T) {
	ribbonLengths := []int{3, 5, 7}
	cut := 13

	if pattern2.MaxRibbonCut(ribbonLengths, cut) != 3 {
		t.Fatal("Exercise 5-3 failed")
	}
}
