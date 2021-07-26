package pattern3

import (
	"math"
)

func Fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func StaircaseRecursive(n int) int {
	dp := make([][]int, 4)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	return staircaseRecursive(n, dp)
}

func staircaseRecursive(n int, dp [][]int) int {
	switch {
	case n <= 0:
		return 0
	case n == 1:
		return 1
	case n == 2:
		return 2
	case n == 3:
		return 4
	}
	stepOne := dp[1][n]
	stepTwo := dp[2][n]
	stepThree := dp[3][n]
	if stepOne == -1 {
		stepOne = staircaseRecursive(n-1, dp)
	}

	if stepTwo == -1 {
		stepTwo = staircaseRecursive(n-2, dp)
		dp[2][n] = stepTwo
	}

	if stepThree == -1 {
		stepThree = staircaseRecursive(n-3, dp)
		dp[3][n] = stepThree
	}

	return stepOne + stepTwo + stepThree
}

func StaircaseTabular(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n]
}

func CountWays(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}

	return countWays(n, dp)
}

func countWays(n int, dp []int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if dp[n] == -1 {
		way1 := countWays(n-1, dp)
		way2 := countWays(n-3, dp)
		way3 := countWays(n-4, dp)

		dp[n] = way1 + way2 + way3
	}

	return dp[n]
}

func CountWaysTabular(n int) int {
	dp := make([]int, n+1)

	if n <= 0 {
		return 0
	}

	if n < 2 {
		return 1
	}

	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	dp[4] = 4
	for i := 5; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-4] + dp[i-1]
	}
	return dp[n]
}

func minReachEnd(steps []int, index int, dp []int) int {
	const NO_VALID_PATH int = -1
	if index+1 == len(steps) {
		return 0
	}

	if index+1 >= len(steps) {
		return NO_VALID_PATH
	}

	if dp[index] != 0 {
		return dp[index]
	}

	if steps[index] == 0 {
		dp[index] = NO_VALID_PATH
	} else {
		way1 := minReachEnd(steps, index+steps[index], dp)
		way2 := minReachEnd(steps, index+1, dp)

		if way2 == NO_VALID_PATH && way1 == NO_VALID_PATH {
			dp[index] = NO_VALID_PATH
		} else if way1 == NO_VALID_PATH {
			dp[index] = way2 + 1
		} else if way2 == NO_VALID_PATH {
			dp[index] = way1 + 1
		} else if way1 < way2 {
			dp[index] = way1 + 1
		} else {
			dp[index] = way2 + 1
		}
	}
	return dp[index]
}

func MinReachEnd(steps []int) int {
	dp := make([]int, len(steps))
	step := minReachEnd(steps, 0, dp)
	return step
}

/*
2,1,1,1,4

i=0
dp
[ 0, 0 , 0, 0, 0]
after
[ 0 1, 1, 0, 0]a

i=1
[1, 2, 2, 0, 0]

i=2
[x1, 2, 1, 3, `]

i=3
*/

func MinReachEndTabular(steps []int) int {
	dp := make([]int, len(steps))
	for i, _ := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for step := 0; step < len(steps); step++ {
		if (step+1) < len(steps) && dp[step]+1 < dp[step+1] {
			dp[step+1] = dp[step] + 1
		}

		if step+steps[step] < len(steps) && dp[step] < dp[step+steps[step]] {
			dp[step+steps[step]] = dp[step] + 1
		}
	}
	return dp[len(steps)-1]
}

/*
1,2,5,2,1,2

[0, inf, inf, , inf, inf]
i=0
[0, 1, 1, 1, inf, inf]
i=1
[0, 1, 1, 1, 3, inf],
i=2
[0, 1, 1, 1, 3, 6]
i=3
[0, 1, 1, 1, 3, 3]
i=4
*/

func MinFeeReachEnd(steps []int) int {
	dp := make([]int, len(steps))
	minSteps := math.MaxInt32

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < len(steps); i++ {
		if (i+1) >= len(steps) && (dp[i]+steps[i]) < minSteps {
			minSteps = steps[i]
		} else if (i+1) < len(steps) && dp[i]+steps[i] < dp[i+1] {
			dp[i+1] = dp[i] + steps[i]
		}
		if (i+2) >= len(steps) && (dp[i]+steps[i]) < minSteps {
			minSteps = steps[i] + dp[i]
		} else if (i+3) < len(steps) && (dp[i] + steps[i]) < dp[i+2] {
			dp[i+2] = dp[i] + steps[i]
		}
		if (i+3) >= len(steps) && (dp[i]+steps[i]) < minSteps {
			minSteps = dp[i] + steps[i]
		} else if (i+3) < len(steps) && (dp[i] + steps[i]) < dp[i+3] {
			dp[i+3] = dp[i] + steps[i]
		}
	}
	return minSteps 
}
