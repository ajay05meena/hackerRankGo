package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	lengthOfString := int64(len(s))
	count := int64(0)
	for _, char := range s{
		if char == 'a' {
			count = count +1
		}
	}

	count = count * n/lengthOfString
	for i, char := range s{
		if char == 'a' && int64(i) < n%lengthOfString{
			count = count + 1
		}
	}
	return count
}


func main() {
	fmt.Println(repeatedString("aba", 10))
}
