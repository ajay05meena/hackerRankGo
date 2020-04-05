package mars_exploration

func MarsExploration(s string) int32 {
	count :=0
	for i, char := range s {
		if i%3 != 1 {
			if string(char) != "S" {
				count ++
			}
		}else {
			if string(char) != "O"{
				count ++
			}
		}
	}
	return int32(count)
}
