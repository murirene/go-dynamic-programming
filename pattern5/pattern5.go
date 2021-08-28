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

/*
    a   b   c   f
    0   0   0   0
b   0   1   1   1
d   0   1   1   1
c   0   1   2   2
f   0   1   2   3

dp[i][j] = max(dp[i-1][j],dp[i][j-1])

if word1[i] == word2[j]
dp[i][j] = dp[i-1][j-1]+1

common := dp[len(word1)-1][len(word2)-1]
return common + len(word1)-common + len(word2) - common

    p   r   o   g   r   a   m   m   i   n   g
d   0   0   0   0   0   0   0   0   0   0   0
y   0   0   0   0   0   0   0   0   0   0   0
n   0   0   0   0   0   0   0   0   0   1   1
a   0   0   0   0   0   1   1   1   1   1   1
m   0   0   0   0   0   1   2   2   2   2   2
i   0   0   0   0   0   1   2   2   3   3   3
c   0   0   0   0   0   1   2   2   3   3   3

3 + 4 + 8 = 15
*/
/*
func ShortestCommonSuperSequence(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))

			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	common := dp[len(word1)][len(word2)]

	return common + len(word1) - common + len(word2) - common
}
*/
func ShortestSuperSequenceRecursive(word1, word2 string, i, j int) int {
	if len(word1) == i {
		return len(word2) - j
	}

	if len(word2) == j {
		return len(word1) - i
	}

	if word1[i] == word2[j] {
		return ShortestSuperSequenceRecursive(word1, word2, i+1, j+1) + 1
	}

	c1 := ShortestSuperSequenceRecursive(word1, word2, i, j+1) + 1
	c2 := ShortestSuperSequenceRecursive(word1, word2, i+1, j) + 1

	return int(math.Min(float64(c1), float64(c2)))
}

func MinimumDeletes(list []int, i, j int) int {
	if len(list) == j {
		return 0
	}

	c1 := math.MaxInt32
	if i == -1 || list[i] < list[j] {
		c1 = MinimumDeletes(list, j, j+1)
	}

	c2 := MinimumDeletes(list, i, j+1) + 1

	return int(math.Min(float64(c1), float64(c2)))
}

func shortestCommonSequence(word1, word2 string, i, j int, dp [][]int) int {
	if len(word2) == j {
		return len(word1) - i
	}

	if len(word1) == i {
		return len(word2) - j
	}

	if dp[i][j] != 0 {
		return dp[i][j]
	}

	if word1[i] == word2[j] {
		return shortestCommonSequence(word1, word2, i+1, j+1, dp) + 1
	}

	cs1 := shortestCommonSequence(word1, word2, i, j+1, dp) + 1
	cs2 := shortestCommonSequence(word1, word2, i+1, j, dp) + 1

	dp[i][j] = int(math.Min(float64(cs1), float64(cs2)))

	return dp[i][j]
}

func ShortestCommonSuperSequence(word1, word2 string) int {
	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
	}
	return shortestCommonSequence(word1, word2, 0, 0, dp)
}

func ShortestSuperCommonSequenceTabular(word1, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}

	for i := range word2 {
		dp[0][i+1] = i + 1
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func MinimumDeletionsSortedSequence(list []int, i, j int) int {
	if len(list) == j {
		return 0
	}

	m := math.MaxInt32
	if i == -1 || list[i] < list[j] {
		m = MinimumDeletionsSortedSequence(list, j, j+1)
	}

	c1 := MinimumDeletionsSortedSequence(list, i, j+1) + 1

	return int(math.Min(float64(m), float64(c1)))
}

func MinimumDeletionsSortedSequenceTabular(list []int) int {
	dp := make([]int, len(list))
	for i := range dp {
		dp[i] = 1
	}

	max := dp[0]
	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[j] < list[i] && dp[j] >= dp[i] {
				dp[i] += 1
				max = int(math.Max(float64(max), float64(dp[i])))
			}
		}
	}

	return len(list) - max
}

/*
         t   o   m   o   r   r   o   w
     0   0   0   0   0   0   0   0   0
t    0   0   0   0   0   0   0   0   0
o    0   0   0   0   1   1   1   1   1
m    0   0   0   0   1   1   1   1   1
o    0   0   0   0   1   1   1   2   2
r    0   0   0   0   1   1   2   2   2
r    0   0   0   0   1   1   2   2   2
o    0   0   0   0   1   1   2   2   2
w    0   0   0   0   1   1   2   2   2


f

for i = 1; i<len(word)
for j=i+1; j<len(word)
dp[i][j] = math.max(dp[i-1][j], dp[i][j-1])

if word[i] == word[j]
dp[i][j] = dp[i-1][j-1]+1

*/

func LongestRepeatingSequence(word string) int {
	dp := make([][]int, len(word)+1)
	for i := range dp {
		dp[i] = make([]int, len(word)+1)
	}

	for i := 1; i <= len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			if word[i-1] == word[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	length := len(word)
	dp[length][length] = int(math.Max(float64(dp[length-1][length]), float64(dp[length][length-1])))

	return dp[length][length]
}

/*
    b   a   x   m   x
a   0   1   1   1   1
x   0   0   1   1   2

    t   o   m   o   r   r   o   w
t   1   1   1   1   1   1   1   1
o   0   1   1   2   2   2   3   3
r   0   0   0   0   2   4   4   4


dp[i][j] = dp[i][j-1]
if word[i] == word[j]
dp[i][j] = dp[i][j] + dp[i-1][j-1]

*/

func SubSequencePatternMatching(word string, pattern string) int {
	dp := make([][]int, len(pattern)+1)
	for i := range dp {
		dp[i] = make([]int, len(word)+1)
	}

	for i := 1; i <= len(pattern); i++ {
		for j := 1; j <= len(word); j++ {
			dp[i][j] = dp[i][j-1]
			if word[j-1] == pattern[i-1] {
				dp[i][j] = int(math.Max(float64(1), float64(dp[i][j-1]+dp[i-1][j-1])))
			}
		}
	}

	return dp[len(pattern)][len(word)]
}

func LongestBitonicSubsequence(list []int, p, i, j int) int {
	if i == len(list) {
		return 0
	}

	if j == len(list) {
		return 0
	}

	c1 := 0

	if p == -1 || (list[p] <= list[i] && list[i] <= list[j]) {
		c1 = LongestBitonicSubsequence(list, i, j, j+1) + 1
	} else if list[i] > list[j] {
		c1 = LongestBitonicSubsequence(list, i, j, j+1) + 1
	}

	c2 := LongestBitonicSubsequence(list, p, i, j+1)

	return int(math.Max(float64(c1), float64(c2)))
}

func LongestBitonicSubsequenceTab(list []int) int {
	max := 1
	dp := make([]int, len(list))

	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(list); i++ {
		p := -1
		for j := 0; j < i; j++ {
			if p == -1 || (list[i] > list[j] && list[j] > list[p]) ||
				(list[i] < list[j]) {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				max = int(math.Max(float64(max), float64(dp[i])))
			}
			p = j
		}
	}

	return max
}

func LongestAlternatingSubsequenceTab(list []int) int {
	max := 1
	dp := make([]int, len(list))

	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(list); i++ {
		p := -1
		for j := 0; j < i; j++ {
			if p == -1 || (list[j] > list[i] && list[j] > list[p]) ||
				(list[j] < list[i] && list[j] < list[p]) {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
				max = int(math.Max(float64(max), float64(dp[i])))
			}
			p = j
		}
	}

	return max
}

func Edit(word1, word2 string) int {
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

	lcs := dp[len(word1)][len(word2)]

	deletes := 0
	inserts := 0
	replace := 0
	if len(word1) > len(word2) {
		deletes = len(word1) - len(word2)
		replace = len(word2) - lcs
	} else {
		inserts = len(word2) - len(word1)
		replace = len(word1) - lcs
	}

	return deletes + replace + inserts
}
