# 18 - Topological Sort (DAG Dependency Resolution)

## 18.1 - Overview & Theoretical Foundations (CLRS 20.4)

* A **Topological Sort** of a directed acyclic graph $G = (V, E)$ is a linear ordering of all its vertices such that if $G$ contains an edge $(u, v)$, then $u$ appears before $v$ in the ordering.
* **CLRS Lemma 20.11:** A directed graph $G$ is acyclic if and only if a depth-first search of $G$ yields **no back edges**.
* **Two Classic Implementations:**
  1. **Kahn's Algorithm (In-Degree BFS):**
     * Maintains an array of in-degrees for each vertex.
     * Starts with all vertices having $\text{in-degree} = 0$.
     * Progressively removes vertices, decrementing neighbors' in-degrees. If the final ordering contains fewer than $|V|$ vertices, a **cycle** is detected.
  2. **DFS with Finishing Times (CLRS 20.4):**
     * Runs DFS to compute finishing times $v.f$ for each vertex.
     * As each vertex is finished, inserts it onto the front of a linked list. The resulting list is a topological sort.

---

## 18.2 - Properties of a problem that suggests Topological Sort

* Determining **build order**, **package compilation**, or **course prerequisite scheduling**.
* Detecting cycles in directed graphs.

---

## 18.3 - Classic Example: Course Schedule II (Kahn's Algorithm)

### Java Implementation

```java
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

public class TopologicalSort {

    public static int[] findOrder(int numCourses, int[][] prerequisites) {
        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) adj.add(new ArrayList<>());
        int[] inDegree = new int[numCourses];

        // Build adjacency graph & in-degree array
        for (int[] prereq : prerequisites) {
            int dest = prereq[0];
            int src = prereq[1];
            adj.get(src).add(dest);
            inDegree[dest]++;
        }

        Queue<Integer> queue = new ArrayDeque<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }

        int[] order = new int[numCourses];
        int index = 0;

        while (!queue.isEmpty()) {
            int u = queue.poll();
            order[index++] = u;

            for (int v : adj.get(u)) {
                inDegree[v]--;
                if (inDegree[v] == 0) {
                    queue.offer(v);
                }
            }
        }

        // If index != numCourses, graph contains a cycle
        return (index == numCourses) ? order : new int[0];
    }
}
```

---

### Go Implementation

```go
package main

// FindOrder finds a topological ordering of courses, or returns empty slice if cyclic
func FindOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		dest, src := p[0], p[1]
		adj[src] = append(adj[src], dest)
		inDegree[dest]++
	}

	queue := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := make([]int, 0, numCourses)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:] // Dequeue
		order = append(order, u)

		for _, v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(order) == numCourses {
		return order
	}
	return []int{} // Cycle detected
}
```

---

## 18.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\Theta(V + E)$ — Constructing the graph takes $\mathcal{O}(E)$. Each vertex enters and leaves the queue once ($\mathcal{O}(V)$), and every edge is inspected once ($\mathcal{O}(E)$).
* **Space Complexity:** $\Theta(V + E)$ to store the adjacency list graph and in-degree table.

---

## 18.5 - Classic LeetCode & CLRS Benchmarks

* **Topological Sort** (CLRS 20.4)
* **Course Schedule I & II** (LeetCode #207, #210)
* **Alien Dictionary** (LeetCode #269)
* **Minimum Height Trees** (LeetCode #310)
* **Sequence Reconstruction** (LeetCode #444)

---

## 18.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Section 20.4: Topological sort (pp. 614–620)
* https://en.wikipedia.org/wiki/Topological_sorting
* https://leetcode.com/problems/course-schedule-ii/
