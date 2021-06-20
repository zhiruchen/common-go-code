package slice

func Filter(s []int, filterFn func(n int) bool) []int {
	result := make([]int, 0, len(s))
	for _, n := range s {
		if filterFn(n) {
			result = append(result, n)
		}
	}

	return result
}

func FilterInPlace(s []int, filterFn func(n int) bool) []int {
	result := s[:0]
	for _, n := range s {
		if filterFn(n) {
			result = append(result, n)
		}
	}

	return result
}
