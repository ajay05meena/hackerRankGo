package save_the_prisoner

func SaveThePrisoner(n int32, m int32, s int32) int32 {
	if (m%n + s-1)%n == 0{
		return n
	}
	return (m%n + s-1)%n

}
