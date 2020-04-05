package camelcase

import "unicode"

func Camelcase(s string) int32 {
	count := 1
	for _, ch := range s{
		if unicode.IsUpper(ch){
			count++
		}
	}
	return int32(count)

}