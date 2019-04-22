package sorting

//Selection - Selection Sort
func Selection(dataSet []int) []int {
	for i := 0; i < len(dataSet); i++ {
		for j := i + 1; j < len(dataSet); j++ {
			if dataSet[j] < dataSet[i] {
				dataSet[i], dataSet[j] = dataSet[j], dataSet[i]
			}
		}
	}
	return dataSet
}
