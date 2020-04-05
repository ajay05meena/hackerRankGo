package picking_numbers

func PickingNumbers(a []int32) int32 {
	aMap := make(map[int32]int32)
	for _, val := range a{
		aMap[val]++
	}
	max := int32(0)
	for k, _ := range aMap{
		tmp1 := aMap[k] + aMap[k-1]
		tmp2 := aMap[k] + aMap[k+1]
		if max < tmp1 {
			max = tmp1
		}
		if max < tmp2 {
			max = tmp2
		}
	}
	return max
}