package integer

import "log"

type Set struct {
	data map[int]struct{}
}

func NewSet(as ...int) *Set {
	s := &Set{
		data: make(map[int]struct{}),
	}
	s.Insert(as...)
	return s
}

func (s *Set) Copy() *Set {
	copys := &Set{
		data: make(map[int]struct{}),
	}
	for k := range s.data {
		copys.data[k] = struct{}{}
	}
	return copys
}

func (s *Set) Contains(a int) bool {
	_, ok := s.data[a]
	return ok
}

func (s *Set) ContainsAll(as ...int) bool {
	for _, a := range as {
		if _, ok := s.data[a]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) ContainsAny(as ...int) bool {
	for _, a := range as {
		if _, ok := s.data[a]; ok {
			return true
		}
	}
	return false
}

func (s *Set) Len() int {
	return len(s.data)
}

func (s *Set) Insert(as ...int) {
	for _, a := range as {
		s.data[a] = struct{}{}
	}
}

func (s *Set) Remove(as ...int) {
	for _, a := range as {
		if _, ok := s.data[a]; !ok {
			log.Fatalln("Nothing to remove when removing:", a)
		}
		delete(s.data, a)
	}
}

func (s *Set) Clear() {
	s.data = make(map[int]struct{})
}
