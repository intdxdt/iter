package iter

import "sort"

//Sort in place slice of ints as unique values
func SortedIntSet(values *[]int) []int {
	var res = MakeUnique(values)
	sort.Ints(res)
	return res
}

//Mkae unique slice of ints
func MakeUnique(values *[]int) []int {
	var dict = make(map[int]struct{})
	var ints = *values
	for _, v := range ints {
		dict[v] = struct{}{}
	}
	var n = len(dict)
	ints = ints[0:n]
	var i = 0
	for k := range dict {
		ints[i] = k
		i += 1
	}
	*values = ints
	return ints
}
