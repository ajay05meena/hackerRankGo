package main

import "fmt"
import "os"

func isValid(str string) string{
	m := make(map[string]int)
	for _, c := range str{
		m[string(c)] = m[string(c)] + 1	
	}
	
	
	occ := make(map[int]int)
	for _, v := range m{
		occ[v] = occ[v] + 1	
	}

	if len(occ) > 2 {
		return "NO"
	}else if len(occ) == 1{
		return "YES"
	}else {
		min := 0
		max := 0
		min_o := 0
		max_o := 0
		for k, v := range occ {
			if min == 0 || min > k{
				min = k
				min_o = v
			}
			if max == 0 || max < k{
				max = k
				max_o = v
			}
		}

		if max - min == 1 && ( max_o == 1 || min_o == 1 ){
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(isValid(os.Args[1]))
}
