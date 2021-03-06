package closest_numbers

import (
	"math"
	"sort"
)

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


func ClosestNumbers(arr []int32) []int32 {
	sort.Sort(foo(arr))
	prev := int32(2147483647)
	min := float64(2147483647)
	for _, n := range arr{
		if min > math.Abs(float64(prev-n)) {
			min = math.Abs(float64(prev-n))
		}
		prev = n
	}

	res := make([] int32, 0)

	prev = int32(2147483647)
	for _, n := range arr{
		if min == math.Abs(float64(prev-n)) {
			res = append(res, prev, n)
		}
		prev = n
	}
	return res
}
