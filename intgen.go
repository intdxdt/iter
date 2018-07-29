package iter

type Igen struct {
	start   int
	current int
}

func (s *Igen) Next() int {
	var cur = s.current
	s.current++
	return cur
}

func NewIgen(start ...int) *Igen {
	var s int
	if len(start) > 0 {
		s = start[0]
	}
	return &Igen{start: s, current: s}
}

//Integer Generator
type Generator struct {
	start, stop, step, v int
	HasNext              bool
	fromValues           bool
	values               []int
}

func NewGenerator(args ...int) *Generator {
	var self = &Generator{}

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

func NewGeneratorOfVals(args ...int) *Generator {
	self := &Generator{}
	self.values = args

	if len(self.values) > 0 {
		self.start = 0
		self.stop = len(self.values)
		self.step = 1
	}

	self.fromValues = true
	self.v = self.start - self.step
	self.updateNext(self.start)
	return self
}

func (gen *Generator) updateNext(v int) {
	if gen.step > 0 {
		gen.HasNext = v < gen.stop
	} else {
		gen.HasNext = v > gen.stop
	}
}

func (gen *Generator) Val() int {
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

func (gen *Generator) Values() []int {
	var vals = make([]int, 0)
	for gen.HasNext {
		vals = append(vals, gen.Val())
	}
	return vals
}

func (gen *Generator) First() (int, bool) {
	if !(gen.HasValues()) {
		return 0, false
	}
	return gen.start, true
}

func (gen *Generator) Last() (int, bool) {
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

func (gen *Generator) HasValues() bool {
	diff := gen.stop - gen.start
	return (diff != 0) && (
		(diff < 0 && gen.step < 0) || (diff > 0 && gen.step > 0))
}
