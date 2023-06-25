package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	//require.Equal(t, 6, maxProduct([]int{2, 3, -2, 4}))
	//require.Equal(t, 0, maxProduct([]int{-2, 0, -1}))
	//require.Equal(t, 4, maxProduct([]int{3, -1, 4}))
	require.Equal(t, 24, maxProduct([]int{2, -5, -2, -4, 3}))
}
