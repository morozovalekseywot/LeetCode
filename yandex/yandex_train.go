package main

import (
	"fmt"
)

func solveB() {
	var n int
	fmt.Scan(&n)
	count, ans := 0, 0
	var v int
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		if v == 1 {
			count++
		} else {
			if count > ans {
				ans = count
			}
			count = 0
		}
	}

	if count > ans {
		ans = count
	}

	fmt.Println(ans)
}

func solveC() {
	var n, last, next int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	fmt.Scan(&last)

	fmt.Println(last)

	for i := 1; i < n; i++ {
		fmt.Scan(&next)
		if next > last {
			fmt.Println(next)
			last = next
		}
	}
}

func solveD() {
	var n int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	gen(n, 0, 0, "")
}

func gen(n, countOpen, countClose int, ans string) {
	if countOpen+countClose == 2*n {
		fmt.Println(ans)
		return
	}

	if countOpen < n {
		gen(n, countOpen+1, countClose, ans+"(")
	}
	if countOpen > countClose {
		gen(n, countOpen, countClose+1, ans+")")
	}
}

func main() {
	solveD()
}
