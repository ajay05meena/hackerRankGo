package reduced_string

import "fmt"

func SuperReducedString(s string)string  {
	out := s
	b := false
	for {
		out, b = reduce(out)
		fmt.Print("  ", out)
		if !b{
			if len(out) == 0{
				return "Empty String"
			}else{
				return out
			}
		}
	}

}

func reduce(s string) (string, bool)  {
	var bs = [] byte (s)
	res := ""
	changed := false
	for i:=0; i < len(bs); {
		if i == len(bs)-1 || bs[i] != bs[i+1]{
			res = res + string(bs[i])
			i++
		}else{
			changed = true
			i = i+2
		}
	}
	return res, changed
}
