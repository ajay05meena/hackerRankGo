package icecream_parlor

func IcecreamParlor(m int32, arr []int32) []int32 {
	dict := make(map[int32] int32)
	for i, n := range arr{
		if dict[m - n] != 0{
			return []int32{dict[m-n], int32(i+1)}
		}else {
			dict[n] = int32(i+1)
		}
	}
	return []int32{0, 0}
}
