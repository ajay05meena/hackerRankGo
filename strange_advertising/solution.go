package strange_advertising

func ViralAdvertising(n int32) int32  {
	liked := 0
	shared := 5
	for n!= 0{
		n--
		liked = liked + shared/2
		shared = shared/2 * 3
	}
	return int32(liked)
}
