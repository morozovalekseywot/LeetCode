package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchRange(t *testing.T) {
	require.Equal(t, []int{3, 4}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	require.Equal(t, []int{-1, -1}, searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
