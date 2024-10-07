package eliotlibs

func GenerateCombinations(A []int, K int) [][]int {
	var result [][]int
	current := make([]int, K)
	var generateCombinationsHelper func(start, currentIndex int)

	generateCombinationsHelper = func(start, currentIndex int) {
		if currentIndex == K {
			combination := make([]int, K)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for i := start; i < len(A); i++ {
			current[currentIndex] = A[i]
			generateCombinationsHelper(i+1, currentIndex+1)
		}
	}

	generateCombinationsHelper(0, 0)
	return result
}
