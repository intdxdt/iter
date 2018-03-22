package iter

import (
	"testing"
	"github.com/franela/goblin"
	"sort"
)

func TestUniq(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("UniqueInts", func() {
		g.It("unique ints", func() {
			vals := []int{4, 6, 8, 9, 6, 7, 5, 6, 4, 3}
			expects := []int{3, 4, 5, 6, 7, 8, 9}
			results := UniqueInts(vals)
			sort.Ints(results)
			g.Assert(results).Equal(expects)
		})
		g.It("sorted unique ints", func() {
			vals := []int{4, 6, 8, 9, 6, 7, 5, 6, 4, 3}
			expects := []int{3, 4, 5, 6, 7, 8, 9}
			results := SortedUniqueInts(vals)
			g.Assert(results).Equal(expects)
		})

	})

}
