package intset

type IntSet struct {
	data []bool
	len  int
}

func New(N int) IntSet {
	return IntSet{
		data: make([]bool, N),
		len:  0,
	}
}

func (s IntSet) Add(elements ...int) {
	for _, element := range elements {
		if !s.data[element] {
			s.data[element] = true
			s.len++
		}
	}
}

func (s IntSet) Remove(elements ...int) {
	for _, element := range elements {
		if s.data[element] {
			s.data[element] = false
			s.len--
		}
	}
}

func (s IntSet) Len() int {
	return s.len
}

func (s IntSet) Contains(element int) bool {
	return s.data[element]
}
