package missing_numbers

import "sort"

type foo [] int32

func (f foo) Len() int {
	return len(f)
}

func (f foo) Swap(i, j int)  {
	f[i], f[j] = f[j], f[i]
}

func (f foo) Less(i, j int) bool  {
	return f[i] < f[j]
}

func MissingNumbers(arr []int32, brr []int32) []int32 {
	sort.Sort(foo(arr))
	sort.Sort(foo(brr))
	res := make([]int32, 0, 0)
	j := 0
	i := 0
	for i < len(arr){
		if j == len(brr){
			res = append(res, arr[i])
			i++
		}else{
			if arr[i] == brr[j] {
				j++
				i++
			}else{
				res = append(res, arr[i])
				i++
			}
 		}
	}
	return res
}
