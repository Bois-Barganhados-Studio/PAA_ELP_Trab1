package bruteforce

func GetSumSets(numbers []int, sum int) []int {
	subset := make([]int, 0)
	var subsetLocal []int
	bitRepresentation := make([]int, len(numbers)+1)

	for bitRepresentation[len(numbers)] == 0 {
		subsetLocal = make([]int, 0)
		for i := 0; i < len(numbers); i++ {
			if bitRepresentation[i] == 1 {
				subsetLocal = append(subsetLocal, numbers[i])
			}
		}
		sumSubset := 0
		for _, number := range subsetLocal {
			sumSubset += number
		}
		if sumSubset == sum {
			subset = subsetLocal
			bitRepresentation[len(numbers)-1] = 1
		}
		// fmt.Println("Bit representation:", bitRepresentation)
		IncrementBitRepresentation(bitRepresentation)
	}
	return subset
}

func IncrementBitRepresentation(bitRepresentation []int) {
	for i := 0; i < len(bitRepresentation); i++ {
		if bitRepresentation[i] == 0 {
			bitRepresentation[i] = 1
			break
		} else {
			bitRepresentation[i] = 0
		}
	}
}
