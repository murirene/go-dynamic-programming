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

/*
prices 2, 6, 7, 10, 13

            0   1   2   3   4   5
1           0   2   4   6   8   10
1,2         0   2   6   8   12  14
1,2,3       0   2   6
1,2,3,4     0
1,2,3,4,5   0

*/

func RodCutOptimalProfit(lengths []int, prices []int, rodLength int) int {
	dp := make([][]int, len(lengths)+1)
	for i := range dp {
		dp[i] = make([]int, rodLength+1)
	}

	for i := 1; i <= len(lengths); i++ {
		for j := 1; j <= rodLength; j++ {
			dp[i][j] = dp[i-1][j]
			if lengths[i-1] > j {
				continue
			}

			rodWithLength := dp[i][j-lengths[i-1]] + prices[i-1]
			if dp[i][j] < rodWithLength {
				dp[i][j] = rodWithLength
			}
		}
	}

	return dp[len(lengths)][rodLength]
}

/*
       0    1   2   3   4   5
       0    0   0   0   0   0
1      0    1   1   1   1   1
1,2    0    1   2   2   3   3
1,2,3  0    1   2   3   4   5
*/

func CoinChangeCount(denominations []int, amount int) int {
	dp := make([][]int, len(denominations)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	for dIndex := 1; dIndex <= len(denominations); dIndex++ {
		for a := 1; a <= amount; a++ {
			dp[dIndex][a] = dp[dIndex-1][a]

			if denominations[dIndex-1] > a {
				continue
			}

			if dp[dIndex][a] < (dp[dIndex][a-denominations[dIndex-1]] + 1) {
				dp[dIndex][a] = dp[dIndex][a-denominations[dIndex-1]] + 1
			}
		}
	}
	return dp[len(denominations)][amount]
}
