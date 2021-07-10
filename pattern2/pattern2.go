package pattern2

// C capacity
// Weights {1, 2, 3}
// Profits {15, 20, 50}

/*
        0    1    2   3   4   5
        0    0    0   0   0   0
1       0    15   30  45  60  75
1, 2    0    15   30  45  60  75
1,2,3   0    15   30  50  65  80   (melon and 2 apples)
*/

func UnboundKnapsack(weights []int, profits []int, capacity int) int {
	dpTable := make([][]int, len(weights)+1)
	for i := range dpTable {
		dpTable[i] = make([]int, capacity+1)
	}

	for i := 1; i < len(dpTable); i++ {
		for j := 1; j < len(dpTable[i]); j++ {
			dpTable[i][j] = dpTable[i-1][j]
			if weights[i-1] > j {
				continue
			}

			if (dpTable[i][j-weights[i-1]] + profits[i-1]) > dpTable[i][j] {
				dpTable[i][j] = dpTable[i][j-weights[i-1]] + profits[i-1]
			}
		}
	}

	return dpTable[len(weights)][capacity]
}
