package iter

import (
	"github.com/franela/goblin"
	"sort"
	"testing"
	"time"
)

func TestUniq(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("UniqueInts", func() {
		g.It("unique ints", func() {
			g.Timeout(1 * time.Hour)
			vals := []int{4, 6, 8, 9, 6, 7, 5, 6, 4, 3}
			expects := []int{3, 4, 5, 6, 7, 8, 9}
			results := UniqueInts(vals)
			sort.Ints(results)
			g.Assert(sort.IntsAreSorted(results)).IsTrue()
			g.Assert(results).Equal(expects)
		})
		g.It("sorted unique ints", func() {
			vals := []int{4, 6, 8, 9, 6, 7, 5, 6, 4, 3}
			expects := []int{3, 4, 5, 6, 7, 8, 9}
			results := SortedIntsSet(vals)
			g.Assert(sort.IntsAreSorted(results)).IsTrue()
			g.Assert(results).Equal(expects)
		})

	})

}
