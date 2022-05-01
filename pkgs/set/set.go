package set

type set struct {
	m map[interface{}]struct{}
}

func newSet() *set {
	return &set{
		m: make(map[interface{}]struct{}),
	}
}

func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(s.m, item)
	}
}

func (s *set) Exists(item interface{}) bool {
	_, exists := s.m[item]
	return exists
}

func (s *set) Size() int {
	return len(s.m)
}
