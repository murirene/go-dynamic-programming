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

/*
abdca
cbda                            count
        a   b   d   c   a
c       0   0   0   1   0           1
b       0   1   0   0   0           2
d       0   0   2   0   0
a       1   0   0   0   1

        p   a   s   s   p   o   r   t
ppsspt  1   1   1   1
*/

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

/*
    lcs = 2
    s1 - lcs = deletes 1
    s2-lcs = inserts 1


  passport
  ppsspt

             p   a  s   s  p   o   r   t
p            1   1  1   1  1   1   1   1   0
p            1   1  1   1  2   2   2   2   0
s            1   1  2   2   2   2  2   2   0
s            1   1  2   3   3   3   3   3   0
p            1   1  2   3   4   4   4   4   0
t            1   1  2   3   4   4   4   5   0


i
p            0   0  0   0   0   0   0   0   0
p            0   0  0   0   0   0   0   0   0
p            0   0  0   0   0   0   0   0   0
p            0   0  0   0   0   0   0   0   0
p            0   0  0   0   0   0   0   0   0
s        0   0   1   0   1   0
s        0   0   1   2   1   0
a        0   0   1   2   2   3
        0   0   0   0   0   0

    5-3 = 2 deletes
    4-3 - 1 insert
    3
*/

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

/*
    4   ,2  ,3  ,6  ,10 ,1  ,12
4   4   4   4   6   10   10  12     4
2   0   2   3   4   5   0   6      5
3   0   0   0   0   0   0   0
6   0   0   0   0   0   0   0

i=0, j=1, +1
    0  1  2  3  4   5  6
    4, 2, 3, 6, 10, 1, 12
    i, 0, j, 1, 4, 2  <
    i: 0, j: 2, 4, 3  <
    i: 0, j: 3, 4, 6  > + 1
    i: 3, j: 4, 6, 10, > + 1
    i: 4, j: 5, 10, 1 <
    i: 4, j: 6, 10, 12 > + 1
    i: 6, j: 7, 12, Out of index return 0 (4)

    if  4 > 2 then skip to 3 until the end is reached. rturn the max

    2, 3, 6, 10, 1, 12
     0, 1,  2, 3, 4
    -4, 10, 3, 7, 15

    i: 0, j: 1 , > + 1
    i: 1, j: 2 , <

2, 3
*/

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
