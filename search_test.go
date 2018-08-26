package iter

import (
	"github.com/franela/goblin"
	"sort"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("sort search", func() {
		g.It("search ints", func() {
			g.Timeout(1 * time.Hour)
			data := []int{3, 4, 5, 6, 7, 8, 9, 11, 20}
			g.Assert(sort.IntsAreSorted(data)).IsTrue()
			g.Assert(SortedSearchInts(data, 0)).IsFalse()
			g.Assert(SortedSearchInts(data, -2)).IsFalse()
			g.Assert(SortedSearchInts(data, 2)).IsFalse()
			g.Assert(SortedSearchInts(data, 1)).IsFalse()
			g.Assert(SortedSearchInts(data, 4)).IsTrue()
			g.Assert(SortedSearchInts(data, 9)).IsTrue()
			g.Assert(SortedSearchInts(data, 11)).IsTrue()
			g.Assert(SortedSearchInts(data, 10)).IsFalse()
			g.Assert(SortedSearchInts(data, 15)).IsFalse()
		})
		g.It("search f64s", func() {
			g.Timeout(1 * time.Hour)
			data := []float64{0.1, 0.2, 0.1 + 0.2, 3, 4, 5, 6, 7, 8, 9, 11, 20}
			g.Assert(sort.Float64sAreSorted(data)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 0.1)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 0.2)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 0.3)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 0)).IsFalse()
			g.Assert(SortedSearchFloat64s(data, -2)).IsFalse()
			g.Assert(SortedSearchFloat64s(data, 2)).IsFalse()
			g.Assert(SortedSearchFloat64s(data, 1)).IsFalse()
			g.Assert(SortedSearchFloat64s(data, 4)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 9)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 11)).IsTrue()
			g.Assert(SortedSearchFloat64s(data, 10)).IsFalse()
			g.Assert(SortedSearchFloat64s(data, 15)).IsFalse()
		})

	})
}
