package leetcode

import (
	"container/heap"
	"math"
)

type Edge struct {
	to int
	w  float64
}

// An EdgeHeap is a min-heap of ints.
type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].w > h[j].w }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://leetcode.com/problems/evaluate-division/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	uniqElements := make(map[string]int, len(equations))
	count := 0
	for _, eq := range equations {
		if _, ok := uniqElements[eq[0]]; !ok {
			uniqElements[eq[0]] = count
			count++
		}
		if _, ok := uniqElements[eq[1]]; !ok {
			uniqElements[eq[1]] = count
			count++
		}
	}

	graph := make([][]Edge, count)
	for i := 0; i < len(equations); i++ {
		idxA, idxB := uniqElements[equations[i][0]], uniqElements[equations[i][1]]
		graph[idxB] = append(graph[idxB], Edge{idxA, values[i]})
		if values[i] != 0.0 {
			graph[idxA] = append(graph[idxA], Edge{idxB, 1.0 / values[i]})
		} else {
			graph[idxA] = append(graph[idxA], Edge{idxB, math.Inf(1)})
		}
	}

	answers := make([]float64, 0, len(queries))
	for _, q := range queries {
		to, ok1 := uniqElements[q[0]]
		from, ok2 := uniqElements[q[1]]
		if !ok1 || !ok2 {
			answers = append(answers, -1.0)
			continue
		}

		// Ищем расстояние с помощью алгоритма Дейкстры
		h := &EdgeHeap{}
		heap.Init(h)

		distance := make([]float64, count)
		for i := 0; i < count; i++ {
			distance[i] = math.Inf(1)
		}

		processed := make([]bool, count)

		distance[from] = 1
		h.Push(Edge{to: from, w: 0})
		for h.Len() > 0 {
			vert := h.Pop().(Edge).to
			if processed[vert] {
				continue
			}
			processed[vert] = true
			for _, edge := range graph[vert] {
				if distance[vert]*edge.w < distance[edge.to] {
					distance[edge.to] = distance[vert] * edge.w
					h.Push(Edge{to: edge.to, w: distance[edge.to]})
				}
			}
		}

		if math.IsInf(distance[to], 1) {
			answers = append(answers, -1)
		} else {
			answers = append(answers, distance[to])
		}
	}

	return answers
}
