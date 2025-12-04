package sorter

import "sort"

type Sorter[T any] struct {
	S  []T
	Fn func(T, T) bool
}

func (s *Sorter[T]) Len() int {
	if s == nil {
		return 0
	}

	return len(s.S)
}

func (s *Sorter[T]) Less(i, j int) bool {
	if s == nil || s.Fn == nil {
		return false
	}

	return s.Fn(s.S[i], s.S[j])
}

func (s *Sorter[T]) Swap(i, j int) {
	if s == nil || len(s.S) == 0 {
		return
	}

	s.S[i], s.S[j] = s.S[j], s.S[i]
}

func Sort[T any](s *Sorter[T]) {
	if s == nil {
		return
	}

	sort.Sort(s)
}
