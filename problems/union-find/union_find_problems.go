package union_find

import (
	"fmt"
	"sort"
)

// --- Union-Find (Disjoint Set Union) data structure ---
type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	r := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x]) // path compression
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px == py {
		return false
	}
	// Union by rank
	if d.rank[px] < d.rank[py] {
		px, py = py, px
	}
	d.parent[py] = px
	if d.rank[px] == d.rank[py] {
		d.rank[px]++
	}
	return true
}

// 1. KruskalMST computes the total weight of the MST (CLRS Chapter 21 / Kruskal)
// edges[i] = {u, v, weight}
func KruskalMST(n int, edges [][3]int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})
	dsu := NewDSU(n)
	totalWeight := 0
	for _, e := range edges {
		if dsu.Union(e[0], e[1]) {
			totalWeight += e[2]
		}
	}
	return totalWeight
}

// 2. CountComponents counts connected components in an undirected graph (LeetCode #323)
func CountComponents(n int, edges [][2]int) int {
	dsu := NewDSU(n)
	for _, e := range edges {
		dsu.Union(e[0], e[1])
	}
	roots := map[int]bool{}
	for i := 0; i < n; i++ {
		roots[dsu.Find(i)] = true
	}
	return len(roots)
}

// 3. FindRedundantConnection returns the edge that forms a cycle (LeetCode #684)
func FindRedundantConnection(edges [][2]int) [2]int {
	n := len(edges)
	dsu := NewDSU(n + 1)
	for _, e := range edges {
		if !dsu.Union(e[0], e[1]) {
			return e
		}
	}
	return [2]int{}
}

// 4. AccountsMerge merges accounts that share an email (LeetCode #721)
func AccountsMerge(accounts [][]string) [][]string {
	// Map each email to a unique ID
	emailToID := map[string]int{}
	emailToName := map[string]string{}
	id := 0
	for _, acc := range accounts {
		name := acc[0]
		for _, email := range acc[1:] {
			if _, exists := emailToID[email]; !exists {
				emailToID[email] = id
				id++
			}
			emailToName[email] = name
		}
	}
	dsu := NewDSU(id)
	for _, acc := range accounts {
		for i := 2; i < len(acc); i++ {
			dsu.Union(emailToID[acc[1]], emailToID[acc[i]])
		}
	}
	// Group emails by root
	groups := map[string][]string{}
	for email, eid := range emailToID {
		root := fmt.Sprintf("%d", dsu.Find(eid))
		groups[root] = append(groups[root], email)
	}
	result := [][]string{}
	for _, emails := range groups {
		sort.Strings(emails)
		name := emailToName[emails[0]]
		merged := append([]string{name}, emails...)
		result = append(result, merged)
	}
	return result
}
