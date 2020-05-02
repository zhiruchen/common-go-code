package mergesort

func MergeSort(vs []int) []int {
	return sort(vs, 0, len(vs))
}

func sort(arr []int, left, right int) []int {
	if (right - 1) == left {
		return []int{arr[0]}
	}

	mid := left + (right-left)/2
	var leftArr []int
	var rightArr []int
	leftArr = append([]int(nil), arr[left:mid]...)
	rightArr = append([]int(nil), arr[mid:right]...)

	arr1 := sort(leftArr, 0, len(leftArr))
	arr2 := sort(rightArr, 0, len(rightArr))
	result := merge(arr1, arr2)
	return result
}

func merge(a1, a2 []int) (result []int) {
	l1, l2 := len(a1), len(a2)

	var i, j int
	for i < l1 && j < l2 {
		if a1[i] <= a2[j] {
			result = append(result, a1[i])
			i++
			continue
		}
		result = append(result, a2[j])
		j++

	}

	for i < l1 {
		result = append(result, a1[i])
		i++
	}

	for j < l2 {
		result = append(result, a2[j])
		j++
	}

	return result
}
