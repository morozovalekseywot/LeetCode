package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPermute(t *testing.T) {
	require.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
}
