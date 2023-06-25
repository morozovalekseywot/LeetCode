package leetcode

import "golang.org/x/exp/constraints"

const (
	INT32_MAX int   = 2147483647
	INT32_MIN int   = -2147483648
	INT64_MAX int64 = 9223372036854775807
	INT64_MIN int64 = -9223372036854775807
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

type Pair struct {
	fr int
	sc int
}

func minInSlice(arr []int) int {
	mn := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < mn {
			mn = arr[i]
		}
	}

	return mn
}

func maxInSlice(arr []int) int {
	mx := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > mx {
			mx = arr[i]
		}
	}

	return mx
}

func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Abs[T constraints.Signed](x T) T {
	if x > 0 {
		return x
	}
	return -x
}
