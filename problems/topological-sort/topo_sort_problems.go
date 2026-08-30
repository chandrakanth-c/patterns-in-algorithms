package topo_sort

import "sort"

// 1. CanFinish determines if all courses can be finished (LeetCode #207)
// Uses Kahn's BFS topological sort to detect cycles.
func CanFinish(numCourses int, prerequisites [][2]int) bool {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegree[pre[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	visited := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited++
		for _, next := range adj[node] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return visited == numCourses
}

// 2. FindOrder returns a valid course ordering (LeetCode #210)
// Kahn's algorithm; returns empty slice if a cycle is detected.
func FindOrder(numCourses int, prerequisites [][2]int) []int {
	inDegree := make([]int, numCourses)
	adj := make([][]int, numCourses)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
		inDegree[pre[0]]++
	}
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	order := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, next := range adj[node] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}

// 3. AlienOrder returns the character ordering in an alien language (LeetCode #269)
// Builds a graph from adjacent word pairs, then topological sort.
func AlienOrder(words []string) string {
	// Build char set and adjacency list
	adj := map[byte][]byte{}
	inDegree := map[byte]int{}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := inDegree[w[i]]; !ok {
				inDegree[w[i]] = 0
			}
		}
	}
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := len(w1)
		if len(w2) < minLen {
			minLen = len(w2)
		}
		// Edge case: w2 is a prefix of w1 — invalid
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}
		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				adj[w1[j]] = append(adj[w1[j]], w2[j])
				inDegree[w2[j]]++
				break
			}
		}
	}
	// Kahn's BFS
	queue := []byte{}
	for ch, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, ch)
		}
	}
	// Sort to ensure deterministic output
	sort.Slice(queue, func(i, j int) bool { return queue[i] < queue[j] })
	result := []byte{}
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		result = append(result, ch)
		for _, next := range adj[ch] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
		sort.Slice(queue, func(i, j int) bool { return queue[i] < queue[j] })
	}
	if len(result) != len(inDegree) {
		return ""
	}
	return string(result)
}

// 4. FindMinHeightTrees returns root labels giving minimum height trees (LeetCode #310)
// Iteratively prune leaf nodes (topological peeling).
func FindMinHeightTrees(n int, edges [][2]int) []int {
	if n == 1 {
		return []int{0}
	}
	adj := make([]map[int]bool, n)
	for i := range adj {
		adj[i] = map[int]bool{}
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = true
		adj[e[1]][e[0]] = true
	}
	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(adj[i]) == 1 {
			leaves = append(leaves, i)
		}
	}
	remaining := n
	for remaining > 2 {
		remaining -= len(leaves)
		newLeaves := []int{}
		for _, leaf := range leaves {
			for neighbor := range adj[leaf] {
				delete(adj[neighbor], leaf)
				if len(adj[neighbor]) == 1 {
					newLeaves = append(newLeaves, neighbor)
				}
			}
		}
		leaves = newLeaves
	}
	return leaves
}
