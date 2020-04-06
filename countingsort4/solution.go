package countingsort4

import (
	"fmt"
	"strconv"
)

func CountSort( arr [] [] string)  {
	//replaceFirstHalfStringsWithDash(arr)
	min, max := 9223372036854775807, -1
	m := make(map[int] string)
	var keys []int
	for i:=0; i< len(arr); i++ {
		if i < len(arr)/2{
			arr[i][1] = "-"
		}

		if min > stringToNum(arr[i][0]){
			min = stringToNum(arr[i][0])
		}
		if max < stringToNum(arr[i][0]){
			max = stringToNum(arr[i][0])
		}

		if m[stringToNum(arr[i][0])] == ""{
			m[stringToNum(arr[i][0])] = arr[i][1]
			keys = append(keys, stringToNum(arr[i][0]))
		}else{
			m[stringToNum(arr[i][0])] =  m[stringToNum(arr[i][0])] + " " + arr[i][1]
		}
	}




	for _, j := range keys{
		if m[j] != ""{
			fmt.Printf(m[j] + " ")
		}
	}
}

func stringToNum(s string) int{
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil{
		panic("not able to parse" + s)
	}
	return int(num)
}

func replaceFirstHalfStringsWithDash(arr [][]string) {
	arrLength := len(arr)
	for i := 0; i < arrLength/2; i++ {
		arr[i][1] = "-"
	}

}
