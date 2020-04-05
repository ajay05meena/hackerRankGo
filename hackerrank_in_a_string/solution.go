package hackerrank_in_a_string

func HackerrankInString(s string) string {
	stringToFind := "hackerrank"
	foundChar := 0
	bs := [] byte (s)
	i :=0
	for _ , ch := range stringToFind {
		for i < len(s){
			if bs[i] == byte(ch) {
				foundChar++
				i++
				break
			}
			i++
		}
	}

	if foundChar == len(stringToFind) {
		return "YES"
	}else {
		return "NO"
	}
}
