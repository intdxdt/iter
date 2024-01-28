package iter

import (
	"github.com/intdxdt/math"
	"sort"
)

// SortedSearchInts - Sorted search of int(x) in a sorted slice
func SortedSearchInts(data []int, x int) bool {
	var bln = false
	var n = len(data) - 1
	var idx = sort.SearchInts(data, x)
	if 0 <= idx && idx <= n {
		bln = x == data[idx]
	}
	return bln
}

// SortedSearchFloat64s - Sorted search of float64s(x) in a sorted slice
func SortedSearchFloat64s(data []float64, x float64) bool {
	var bln = false
	var n = len(data) - 1
	var idx = sort.SearchFloat64s(data, x)
	if 0 <= idx && idx <= n {
		bln = math.FloatEqual(x, data[idx])
	}
	return bln
}
