package main

var suspicious []bool
var graph [][]int

func remainingMethods(n int, k int, invocations [][]int) []int {
	suspicious = make([]bool, n)
	graph = make([][]int, n)

	for _, edge := range invocations {
		from := edge[0]
		to := edge[1]

		graph[from] = append(graph[from], to)
	}

	dfsRMFP(k)

	for _, edge := range invocations {
		from := edge[0]
		to := edge[1]

		if !suspicious[from] && suspicious[to] {
			ans := make([]int, n)

			for i := 0; i < n; i++ {
				ans[i] = i
			}

			return ans
		}
	}

	var ans []int

	for i := 0; i < n; i++ {
		if !suspicious[i] {
			ans = append(ans, i)
		}
	}

	return ans
}

func dfsRMFP(node int) {
	if suspicious[node] {
		return
	}

	suspicious[node] = true

	for _, next := range graph[node] {
		dfsRMFP(next)
	}
}
