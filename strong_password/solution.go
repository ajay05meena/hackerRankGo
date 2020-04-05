package strong_password

import "unicode"

func MinimumNumber(n int32, password string) int32 {
	m:= int(6)
	digit, lowercase, uppercase, spclcase := false, false, false, false
	for _, ch := range password {
		if unicode.IsDigit(ch) {
			digit = true
		}else if unicode.IsLower(ch){
			lowercase = true
		} else if unicode.IsUpper(ch){
			uppercase = true
		}else {
			spclcase = true
		}
	}
	count:=0
	if !digit {m--; count++}
	if !lowercase {m--; count++}
	if !uppercase {m--; count++}
	if !spclcase {m--; count++}

	if len(password) > m {
		return int32(count)
	}else {
		return int32(count + m - len(password))
	}


}
