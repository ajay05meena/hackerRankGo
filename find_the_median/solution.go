package find_the_median

import "sort"

type foo [] int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f foo) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}


func FindMedian(arr []int32) int32 {
	sort.Sort(foo(arr))
	return arr[len(arr)/2]
}
