package main

import "fmt"

func pw(str string) map[int]bool {
	m := make(map[int]bool)
	cw := 0
	var prev int
	for i, c := range str{ 
		if i == 0 {
			prev = int(c)
			cw = int(c) - 96
		}else{
			m[cw] = true
			if prev == int(c) {
				cw = cw + int(c)-96 
			}else{
				prev = int(c)
				cw = int(c) - 96
			}
		}
	}
	m[cw] = true
	return m
}

func main() {
	m := pw("abccddde")
	for k, _ := range m{
		fmt.Println(k)
	}

}
