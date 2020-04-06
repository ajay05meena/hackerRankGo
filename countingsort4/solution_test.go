package countingsort4_test

import (
	"hackerRankGo/countingsort4"
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := [][]string {{"0", "a"}, {"1", "b"}, {"0" , "c"}, {"1", "d"}}
	countingsort4.CountSort(arr)
	print(arr)
}
