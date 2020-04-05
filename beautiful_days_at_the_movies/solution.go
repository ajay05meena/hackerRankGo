package beautiful_days_at_the_movies

import "math"

func BeautifulDays(start, end, divisor int32) int{
	count := 0
	for start <= end{
		rev := reverse(start, 0)
		if int32(math.Abs(float64(rev-start))) % divisor == 0 {
			count++
		}
		start++
	}
	return count
}

func reverse(num, output int32) int32{
	if num != 0{
		output = output * 10
		output = output + num%10
		num = num/10
		return reverse(num, output)
	}else{
		return output
	}
}
