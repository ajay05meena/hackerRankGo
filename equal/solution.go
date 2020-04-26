package equal

func Equal(arr [] int32) int32 {
	res := int32(147483647)
	min := min(arr)
	for i:=0; i<5; i++{
		currentRun := int32(0)
		target := min - int32(i)
		if target < 0 {
			continue
		}
		for _, ele := range arr{
			currentRun += (ele - target)/5 + (ele - target)%5/2 + (ele - target)%5%2
		}
		if currentRun < res{
			res = currentRun
		}
	}
	return res
}

func min(arr []int32) int32 {
	min := int32(147483647)
	for _, v := range arr{
		if v < min{
			min = v
		}
	}
	return min
}
