package sorting

func merge(arr1 []int, arr2 []int) []int {
	n, m := len(arr1), len(arr2)

	var i, j int

	res := make([]int, 0, n+m)

	for i < n && j < m {
		if arr1[i] <= arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}

	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)

	return res
}

func MergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	mid := n / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)

}
