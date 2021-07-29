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
    for i:=0; i<len(dp);i++{
        dp[i] = make([]int, len(word))
        dp[i][i] = 1
    }

    for i:=len(word);i>=0;i-- {
        for j:=i+1;j<len(word);j++ {
            if word[i] == word[j] {
                size1 := dp[i+1][j-1] + 2 // palindrome include the 2
                size2 := dp[i][j-1] // palindrome without including it
                dp[i][j] =int( math.Max(float64(size1),float64(size2))) 
            } else {
                dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
            }
        }
    }

    return dp[0][len(word)-1]
}
