package dputil

func GetSum(list []int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum 
}

func MakeDpBoolTable(rowSize int, columnSize int) [][]bool {
	dpTable := make([][]bool, rowSize)
	for i := 0; i < rowSize; i++ {
		dpTable[i] = make([]bool, columnSize)
	}

	return dpTable
}

