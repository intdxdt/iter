package iter

import (
	"github.com/franela/goblin"
	"testing"
)

func TestGenerator(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Generator", func() {
		rng := []int{4, 6, 8, 9, 6, 7}
		var val int
		var ok bool
		g.It("int iter range", func() {

			var gen = NewGenerator(10)
			val, ok = gen.First()
			g.Assert(ok).IsTrue()
			g.Assert(val).Equal(0)

			val, ok = gen.Last()
			g.Assert(val).Equal(9)
			g.Assert(ok).IsTrue()

			for i := 0; gen.HasNext && i < 10; i++ {
				g.Assert(gen.Val()).Equal(i)
			}

			gen = NewRange(10)
			val, ok = gen.First()
			g.Assert(ok).IsTrue()
			g.Assert(val).Equal(0)

			val, ok = gen.Last()
			g.Assert(val).Equal(9)
			g.Assert(ok).IsTrue()

			for i := 0; gen.HasNext && i < 10; i++ {
				g.Assert(gen.Val()).Equal(i)
			}

			gen = NewRange(0, 10)
			val, ok = gen.First()
			g.Assert(ok).IsTrue()
			g.Assert(val).Equal(0)

			val, ok = gen.Last()
			g.Assert(val).Equal(9)
			g.Assert(ok).IsTrue()
			for i := 0; gen.HasNext && i < 10; i++ {
				g.Assert(gen.Val()).Equal(i)
			}

			gen = NewRange(1, 10, 2)
			val, ok = gen.First()
			g.Assert(ok).IsTrue()
			g.Assert(val).Equal(1)

			val, ok = gen.Last()
			g.Assert(val).Equal(9)
			g.Assert(ok).IsTrue()

			for i := 1; gen.HasNext && i < 10; i += 2 {
				g.Assert(gen.Val()).Equal(i)
			}

			gen = NewGenerator(1, 10, 2)
			g.Assert(gen.Values()).Eql([]int{1, 3, 5, 7, 9})

			gen = NewGenerator(10, 0, -3)
			val, ok = gen.First()
			g.Assert(ok).IsTrue()
			g.Assert(val).Equal(10)

			val, ok = gen.Last()
			g.Assert(val).Equal(1)
			g.Assert(ok).IsTrue()

			for i := 10; gen.HasNext && i > 0; i -= 3 {
				g.Assert(gen.Val()).Equal(i)
			}

			gen = NewGenerator(10, 0, 3)
			g.Assert(gen.Values()).Eql([]int{})
			gen = NewGenerator(0, 10, -3)
			g.Assert(gen.Values()).Eql([]int{})

			g.Assert(gen.HasValues()).IsFalse()

			val, ok = gen.First()
			g.Assert(ok).IsFalse()
			g.Assert(val).Equal(0)

			val, ok = gen.Last()
			g.Assert(val).Equal(0)
			g.Assert(ok).IsFalse()

		})

		g.It("int range as values", func() {
			var values = Range[int]()
			g.Assert(len(values) == 0).IsTrue()
			var o1 = Range[uint](10)
			g.Assert(len(o1) == 10).IsTrue()
			var o2 = Range[uint](1, 10)
			g.Assert(len(o2) == 9).IsTrue()
			g.Assert([]uint{1, 2, 3, 4, 5, 6, 7, 8, 9}).Equal(o2)
			var o3 = Range[uint](1, 10, 2)
			g.Assert([]uint{1, 3, 5, 7, 9}).Equal(o3)
			var o4 = Range[uint](1, 10, 3)
			g.Assert([]uint{1, 4, 7}).Equal(o4)
			var o5 = Range[int64](10, 1, -3)
			g.Assert([]int64{10, 7, 4}).Equal(o5)
		})

		g.It("int iter as values", func() {
			var gen = NewGeneratorOfVals(rng...)
			for i := 0; gen.HasNext; i++ {
				g.Assert(gen.Val()).Equal(rng[i])
			}
			gen = NewRangeOfVals(rng...)
			for i := 0; gen.HasNext; i++ {
				g.Assert(gen.Val()).Equal(rng[i])
			}
		})
		g.It("int iter as values should panic", func() {
			var gen = NewGeneratorOfVals([]int{0, 1, 2}...)
			defer func() {
				r := recover()
				g.Assert(r != nil).IsTrue()
			}()
			for i := 0; i < 10; i++ {
				g.Assert(gen.Val()).Equal(i)
			}
		})
	})

}

func TestIntGen(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Igen", func() {
		g.It("int generator", func() {
			rg := NewIgen(-5)
			var integers []int
			g.Assert(rg.start).Equal(-5)
			for i := 0; i < 11; i++ {
				integers = append(integers, rg.Next())
			}
			g.Assert(rg.current).Equal(6)
			g.Assert(integers).Equal([]int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5})
		})
	})
}
