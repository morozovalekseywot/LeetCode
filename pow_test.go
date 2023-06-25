package leetcode

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMyPow(t *testing.T) {
	require.Equal(t, float64(512), myPow(2, 9))
}
