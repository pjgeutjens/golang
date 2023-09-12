package bubblesort

func BubbleSort(data []int) []int {
	for j := len(data); j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if data[i] > data[i+1] {
				swap(data, i, i+1)
			}
		}
	}
	return data
}

func swap(data []int, i, j int) {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}
