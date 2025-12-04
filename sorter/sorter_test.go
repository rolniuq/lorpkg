package sorter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	s := &Sorter[int]{
		S: []int{5, 3, 4, 2, 1},
		Fn: func(i1, i2 int) bool {
			return i1 < i2
		},
	}

	Sort(s)
	require.Equal(t, s.S, []int{1, 2, 3, 4, 5})
}
