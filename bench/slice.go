package bench

func AppendToNOMakeSlice(count int32) {
	var slice []int32
	var idx int32 = 0
	for ; idx < count; idx++ {
		slice = append(slice, idx)
	}
}

func AppendToMakeSlice(count int32) {
	var slice = make([]int32, count)
	var idx int32 = 0
	for ; idx < count; idx++ {
		slice[idx] = idx
	}
}

func AppendToMakeSliceWithCap(count int32) {
	var slice = make([]int32, 0, count)
	var idx int32 = 0
	for ; idx < count; idx++ {
		slice = append(slice, idx)
	}
}
