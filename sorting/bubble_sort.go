package sorting

func BubbleSort(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func BubbleSortOptimized(arr []int) {
	n := len(arr)
	swapped := true
	for i := 0; i < n && swapped; i++ {
		swapped = false
		for j := 0; j < n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}
}
