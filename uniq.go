package iter

import "sort"

// Sort in place slice of ints as unique values
func SortedIntsSet(values []int) []int {
	var set = UniqueInts(values)
	sort.Ints(set)
	return set
}

// Mkae unique slice of ints
func UniqueInts(values []int) []int {
	var dict = make(map[int]struct{}, len(values))
	for _, v := range values {
		dict[v] = struct{}{}
	}
	var set = make([]int, len(dict))
	var i = 0
	for k := range dict {
		set[i] = k
		i += 1
	}
	return set
}
