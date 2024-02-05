package iter

import (
	"golang.org/x/exp/constraints"
	"sync"
)

type Igen[T constraints.Signed] struct {
	sync.RWMutex
	start   T
	current T
}

func (s *Igen[T]) Next() T {
	s.Lock()
	var cur = s.current
	s.current++
	s.Unlock()
	return cur
}

func NewInfRange[T constraints.Signed](start ...T) *Igen[T] {
	return NewIgen[T](start...)
}

func NewIgen[T constraints.Signed](start ...T) *Igen[T] {
	var s T
	if len(start) > 0 {
		s = start[0]
	}
	return &Igen[T]{start: s, current: s}
}

// Generator - integer generator
type Generator[T constraints.Signed] struct {
	start, stop, step, v T
	HasNext              bool
	fromValues           bool
	values               []T
}

func NewRange[T constraints.Signed](args ...T) *Generator[T] {
	return NewGenerator[T](args...)
}

func NewGenerator[T constraints.Signed](args ...T) *Generator[T] {
	var self = &Generator[T]{}
	if len(args) == 1 {
		self.start, self.stop, self.step = 0, args[0], 1
	} else if len(args) == 2 {
		self.start, self.stop, self.step = args[0], args[1], 1
	} else if len(args) == 3 {
		self.start, self.stop, self.step = args[0], args[1], args[2]
	}

	self.v = self.start - self.step
	self.updateNext(self.start)
	return self
}

func Range[T constraints.Integer](args ...T) []T {
	var start, stop, step T
	if len(args) == 1 {
		start, stop, step = 0, args[0], 1
	} else if len(args) == 2 {
		start, stop, step = args[0], args[1], 1
	} else if len(args) == 3 {
		start, stop, step = args[0], args[1], args[2]
	}

	var results = make([]T, 0, 32)

	if step > 0 {
		for i := start; i < stop; i += step {
			results = append(results, i)
		}
	} else {
		for i := start; i > stop; i += step {
			results = append(results, i)
		}
	}

	return results
}

func NewRangeOfVals[T constraints.Signed](args ...T) *Generator[T] {
	return NewGeneratorOfVals[T](args...)
}

func NewGeneratorOfVals[T constraints.Signed](args ...T) *Generator[T] {
	self := &Generator[T]{}
	self.values = args

	if len(self.values) > 0 {
		self.start = 0
		self.stop = T(len(self.values))
		self.step = 1
	}

	self.fromValues = true
	self.v = self.start - self.step
	self.updateNext(self.start)
	return self
}

func (gen *Generator[T]) updateNext(v T) {
	if gen.step > 0 {
		gen.HasNext = v < gen.stop
	} else {
		gen.HasNext = v > gen.stop
	}
}

func (gen *Generator[T]) Val() T {
	gen.v += gen.step

	if (gen.step > 0 && gen.v >= gen.stop) ||
		(gen.step < 0 && gen.v <= gen.stop) {
		panic("generator out of range")
	}

	gen.updateNext(gen.v + gen.step)
	if gen.fromValues {
		return gen.values[gen.v]
	}
	return gen.v
}

func (gen *Generator[T]) Values() []T {
	var vals = make([]T, 0)
	for gen.HasNext {
		vals = append(vals, gen.Val())
	}
	return vals
}

func (gen *Generator[T]) First() (T, bool) {
	if !(gen.HasValues()) {
		return 0, false
	}
	return gen.start, true
}

func (gen *Generator[T]) Last() (T, bool) {
	if !(gen.HasValues()) {
		return 0, false
	}

	size := (gen.stop - gen.start) / gen.step
	end := gen.start + (size * gen.step)
	if end == gen.stop {
		end = end - gen.step
	}
	return end, true
}

func (gen *Generator[T]) HasValues() bool {
	diff := gen.stop - gen.start
	return (diff != 0) && ((diff < 0 && gen.step < 0) || (diff > 0 && gen.step > 0))
}
