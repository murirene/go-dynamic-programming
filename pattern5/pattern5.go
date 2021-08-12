package pattern5

import (
	"fmt"
	"math"
)

func lcSubstringHelper(word1, word2 string, i, j, count int, dp map[string]int) int {
	if i == len(word1) || j == len(word2) {
		return count
	}

	key := fmt.Sprintf("%d|%d|%d", i, j, count)

	if v, found := dp[key]; found {
		return v
	}

	if word1[i] == word2[j] {
		count = lcSubstringHelper(word1, word2, i+1, j+1, count+1, dp)
	}

	sum := int(math.Max(float64(lcSubstringHelper(word1, word2, i+1, j, 0, dp)), float64(lcSubstringHelper(word1, word2, i, j+1, 0, dp))))
	dp[key] = int(math.Max(float64(sum), float64(count)))

	return dp[key]
}

func LcSubstring(word1, word2 string) int {
	dp := make(map[string]int)
	return lcSubstringHelper(word1, word2, 0, 0, 0, dp)
}

func LcSubstringTabular(word1, word2 string) int {
	count := 0
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				count = int(math.Max(float64(count), float64(dp[i][j])))
			}
		}
	}

	return count
}

func LcSubsequence(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func LcMinimalEdits(word1, word2 string) (int, int) {
	inserts := 0
	deletes := 0

	// least common subsequence
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	common := dp[len(word1)][len(word2)]
	inserts = len(word1) - common
	deletes = len(word2) - common
	return inserts, deletes
}

func LisHelper(list []int, i, j int) int {
	if j == len(list) {
		return 0
	}

	include := 0
	if i == -1 || list[i] < list[j] {
		include = LisHelper(list, j, j+1) + 1
	}

	exclude := LisHelper(list, i, j+1)
	return int(math.Max(float64(include), float64(exclude)))
}

func Lis(list []int) int {
	return LisHelper(list, -1, 0)
}

func LisTabular(list []int) int {
	dp := make([]int, len(list))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[j] < list[i] && dp[j] == dp[i] {
				dp[i] += 1
			}
		}
	}

	return dp[len(list)-1]
}

func LongestIncreasingSequence(list []int) int {
	dp := make([]int, len(list))
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] > list[j] && dp[i] <= dp[j] {
				dp[i] += 1
			}
		}
	}

	return dp[len(list)-1]
}

func LongestSumSequence(list []int) int {
	dp := make([]int, len(list))
	for i := range dp {
		dp[i] = list[i]
	}

	for i := 1; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] > list[j] && (dp[i]-list[i]) < dp[j] {
				dp[i] = list[i] + dp[j]
			}
		}
	}

	return dp[len(list)-1]
}
