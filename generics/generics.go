package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}
	last := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return last, true
}

func(s *Stack[T]) Push(new T) {
	s.values = append(s.values, new)
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.values) == 0
}