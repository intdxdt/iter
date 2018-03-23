package iter

import (
	"testing"
	"github.com/franela/goblin"
	"time"
)

func TestIntArrayString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Int Array String", func() {
		g.It("int string", func() {
			g.Timeout(1 * time.Hour)
			vals := []int{6, 4, 3}
			g.Assert(IntArrayString(vals)).Equal("[6, 4, 3]")
		})
	})
}
