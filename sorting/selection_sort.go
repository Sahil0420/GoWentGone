package sorting

func Selection_Sort(arr []int) {
	count := len(arr)
	if count == 1 {
		return
	}
	for i := 0; i < count; i++ {
		minIndex := i

		for j := i + 1; j < count; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
