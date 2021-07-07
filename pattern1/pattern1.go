package pattern1

import "github.com/murirene/go-dynamic-programming/dputil"

// Given a set of positive numbers, determine if there exists a subset whose sum is equal to a given number ‘S’.
func HasSubsetSum(target int, list []int) bool {
	dpTable := dputil.MakeDpBoolTable(len(list)+1, target+1)
	for i := 1; i <= len(list); i++ {
		for j := 1; j <= target; j++ {
			dpTable[i][j] = dpTable[i-1][j]
			if list[i-1] > j {
				continue
			}

			if list[i-1] == j {
				dpTable[i][j] = true
				continue
			}

			sumIndex := j - list[i-1]

			if dpTable[i-1][sumIndex] == true {
				dpTable[i][j] = true
			}
		}
	}

	return dpTable[len(list)][target]
}

/*
   Given a set of positive numbers, find if we can partition it into two subsets such that the sum of elements in both the subsets is equal.
   given
*/
func HasSumPair(list []int) bool {
	target := dputil.GetSum(list)

	if target%2 != 0 {
		return false
	}

	target = target / 2

	return HasSubsetSum(target, list)
}

/*
   Given a set of positive numbers, partition the set into two subsets with a minimum difference between their subset sums.
*/
func HasSumPairMinimumDistance(distance int, list []int) bool {
	target := dputil.GetSum(list)
	if target%2 != 0 {
		target = target - distance
		if target%2 != 0 {
			return false
		}
	}
	target = target / 2

	return HasSubsetSum(target, list)
}

// Given a set of positive numbers, find the total number of subsets whose sum is equal to a given number ‘S’.
/*

    1, 1, 2, 3
    4
           0 1 2 3 4
           0 0 0 0 0
1          0 1 0 0 0
1,1        0 2 1 0 0
1,1,2      0 2 2 2 1
1,1,2, 3   0 2 2 3 3

*/

func SubsetCount(target int, list []int) int {
	dpTable := dputil.MakeDpIntTable(len(list)+1, target+1)
	for i := 1; i <= len(list); i++ {
		for j := 1; j <= target; j++ {
			dpTable[i][j] = dpTable[i-1][j]
			if list[i-1] > j {
				continue
			}

			if list[i-1] == j {
				dpTable[i][j] = dpTable[i][j] + 1
				continue
			}

			partIndex := j - list[i-1]
			if dpTable[i-1][partIndex] > 0 {
				dpTable[i][j] = dpTable[i-1][partIndex] + dpTable[i][j]
			}
		}
	}

	return dpTable[len(list)][target]
}
