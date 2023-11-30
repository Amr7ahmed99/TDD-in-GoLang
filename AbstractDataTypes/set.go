package adt

type Set struct {
	values []string
}

func NewSet() *Set {
	return &Set{values: make([]string, 0)}
}

func (s *Set) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Set) Add(item string) {
	s.values = append(s.values, item)
}

func (s *Set) Size() int {
	return len(s.values)
}

func (s *Set) Contains(item string) bool {
	for _, val := range s.values {
		if val == item {
			return true
		}
	}
	return false
}
