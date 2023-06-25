package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMin(t *testing.T) {
	require.Equal(t, 1, findMin([]int{10, 1, 10, 10, 10}))
}
