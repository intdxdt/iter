package iter

import "sort"

func SortedUniqueInts(values []int) []int {
	var res = UniqueInts(values)
	sort.Ints(res)
	return res
}

func UniqueInts(values []int) []int {
	var dict = make(map[int]struct{})
	for _, v := range values {
		dict[v] = struct{}{}
	}
	var res = make([]int, 0, len(dict))
	for k := range dict {
		res = append(res, k)
	}
	return res
}
