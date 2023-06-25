package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	require.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	require.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
	require.Equal(t, -101, threeSumClosest([]int{-100, -98, -2, -1}, -101))
	nums := []int{2, 3, 8, 9, 10}
	require.Equal(t, 15, threeSumClosest(nums, 16))
}
