package pattern1

/*
   Given a set of positive numbers, find if we can partition it into two subsets such that the sum of elements in both the subsets is equal.
   given
*/

func getTarget(list []int) int {
	target := 0

	for v := range list {
		target += v
	}

	if target%2 != 0 {
		return 0
	}

	return target / 2
}

func makeDpTable(rowSize int, columnSize int) [][]int {
	dpTable := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		dpTable[i] = make([]int, columnSize)
	}

	return dpTable
}

func HasSumPair(list []int) bool {
	target := getTarget(list)
	dpTable := makeDpTable(len(list)+1, target+1)

	for i := 1; i <= len(list); i++ {
		for j := 1; j <= target; j++ {
			dpTable[i][j] = dpTable[i-1][j]
			if list[i-1] > j {
				continue
			}

			if list[i-1] == j {
				dpTable[i][j] += 1
				continue
			}

			sumIndex := j - list[i-1]

			if dpTable[i-1][sumIndex] > 0 {
				dpTable[i][j] += 1
			}
		}
	}

	return dpTable[len(list)][target] > 1
}
