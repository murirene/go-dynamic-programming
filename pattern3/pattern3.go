package pattern3

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
