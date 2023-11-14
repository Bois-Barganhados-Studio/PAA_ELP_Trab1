package backtracking

func GetSumSets(numbers []int, sum int, subSet []int, index int) []int {
	//fmt.Println("Chamada GetSumSets(", numbers, ",", sum, ",", subSet, ",", index, ")")
	if sumAll(subSet) == sum {
		return subSet
	}
	originalSubSet := subSet
	// fmt.Println("Testando Subset:", subSet)
	for i := index; i < len(numbers); i++ {
		// fmt.Println("Add number:", numbers[i])
		subSet = append(subSet, numbers[i])
		if sumAll(subSet) == sum {
			// fmt.Println("Subset is sum:", subSet)
			return subSet
		} else if sumAll(subSet) > sum {
			// fmt.Println("Subset is greater than sum:", subSet, " Removing last element.")
			subSet = subSet[:len(subSet)-1]
			continue
		} else {
			// fmt.Println("Subset is less than sum:", subSet, " index is:", i+1)
			subSet = GetSumSets(numbers, sum, subSet, i+1)
			if sumAll(subSet) == sum {
				return subSet
			}
		}
	}
	if len(originalSubSet) > 0 {
		return originalSubSet[:len(originalSubSet)-1]
	} else {
		return nil
	}
}

func sumAll(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
