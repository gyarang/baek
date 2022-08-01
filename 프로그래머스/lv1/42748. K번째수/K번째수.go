import "sort"

func solution(array []int, commands [][]int) []int {
	result := make([]int, len(commands))

	for i := 0; i < len(commands); i++ {
		subArr := append(array[:0:0], array[commands[i][0]-1: commands[i][1]]...)
		sort.Ints(subArr)
		result[i] = subArr[commands[i][2]-1]
	}

	return result
}