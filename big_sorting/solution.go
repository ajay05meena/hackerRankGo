package big_sorting

import "sort"

type numStringSort [] string

func (s numStringSort) Len() int {
	return len(s)
}

func (s numStringSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s numStringSort) Less(i, j int) bool {
	 if len(s[i]) == len(s[j]){
	 	 s1 := []byte(s[i])
	 	 s2 := []byte(s[j])
	 	 n := len(s[i])
	 	 k := 0
	 	 for k<n {
	 	 	if s1[k] == s2[k]{
	 	 		k++
			}else{
				return s1[k] < s2[k]
			}
		 }
	 }else{
	 	return len(s[i]) < len(s[j])
	 }
	 return true
}

func BigSorting(unsorted []string) []string {
	sort.Sort(numStringSort(unsorted))
	return unsorted
}
