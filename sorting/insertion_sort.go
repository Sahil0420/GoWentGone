package sorting

func Insertion_Sort(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}
	var key int
	for i := 1; i < n-1; i++ {
		key = arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
