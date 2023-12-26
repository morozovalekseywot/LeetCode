package leetcode

func getSimpleMod(i, n int) int {
	if i < n {
		return i
	}

	return i - n
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	start, tank, cur := 0, 0, 0
	for start < n {
		tank += gas[cur]
		if tank < cost[cur] {
			if cur < start {
				break
			}
			start = cur + 1
			cur = start
			tank = 0
		} else {
			next := getSimpleMod(cur+1, n)
			if next == start {
				return start
			}

			tank -= cost[cur]
			cur = next
		}
	}

	return -1
}
