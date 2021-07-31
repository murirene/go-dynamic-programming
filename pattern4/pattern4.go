package pattern4

import (
	"math"
)

/*
        abdbca
                a b d b c a
                1 1 1 3 3 5
b,d,b,c,a       0 1 1 3 3 3
d,b,c,a         0 0 1 1 1 1
b,c,a           0 0 0 1 1 1
a,c             0 0 0 0 1 1
a               0 0 0 0 0 1


if a[i] == a[j] m[i][j] = math(m[i+1][j-1j])+2 , m[i-1][j-1] )
if a[i] != a[j] m[i][j] = math(m[i-1][j],m[i][m[j-1])


        c d d p d
c       1 1 2 2 3
d,d,p,d 0 1 2 2 3
d,p,d   0 0 1 1 3
p,d     0 0 0 1 1
d       0 0 0 0 1
*/

func LpsSubsequence(word string) int {
	dp := make([][]int, len(word))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word))
		dp[i][i] = 1
	}

	for i := len(word); i >= 0; i-- {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] {
				size1 := dp[i+1][j-1] + 2 // palindrome include the 2
				size2 := dp[i][j-1]       // palindrome without including it
				dp[i][j] = int(math.Max(float64(size1), float64(size2)))
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return dp[0][len(word)-1]
}

/*
a b d b c a

          a   b   d   b   c   a
abdbca    1   1   1   3   1   1
bdbca     0   1   1   3   1   1
dbca      0   0   1   1   1   1
bca       0   0   0   1   1   1
ca        0   0   0   0   1   1
a         0   0   0   0   0   1


cddpd

         c   d   d   p   d
cddpd    1   1   2   2   3
ddpd     0   1   2   2   3
dpd      0   0   1   1   3
pd       0   0   0   1   1
d        0   0   0   0   1

c d d p d
c d d p d
c 0, 1
    c d d p d
        /\
ddpd        cddp
/\

*/

func lpsSubstringHelper(word string, i int, j int) int {
	if i > j {
		return 0
	}

	if i == j {
		return 1
	}

	if word[i] == word[j] {
		palindrome := lpsSubstringHelper(word, i+1, j-1)
		if (j - i - 1) == palindrome {
			return palindrome + 2
		}
	}

	palindromeLeft := lpsSubstringHelper(word, i+1, j)
	palindromeRight := lpsSubstringHelper(word, i, j-1)

	return int(math.Max(float64(palindromeLeft), float64(palindromeRight)))
}

func LpsSubstringRecursive(word string) int {
	return lpsSubstringHelper(word, 0, len(word)-1)
}
