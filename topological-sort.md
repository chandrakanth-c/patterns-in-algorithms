# 18 - Topological Sort (DAG Dependency Resolution)

## 18.1 - Overview

* **Topological Sorting** provides a linear ordering of vertices in a **Directed Acyclic Graph (DAG)** such that for every directed edge $u \to v$, vertex $u$ comes before vertex $v$ in the ordering.
* If the graph contains a **cycle**, a valid topological ordering is impossible.
* Two primary algorithms:
  1. **Kahn's Algorithm (In-Degree BFS):** Track incoming edges per vertex (`inDegree`). Enqueue nodes with `inDegree == 0`.
  2. **DFS Post-Order Reversal:** Perform DFS and push completed nodes to a stack/list, detecting cycles via node coloring (Unvisited, Visiting, Visited).

---

## 18.2 - Properties of a problem that suggests Topological Sort

* Problem involves **prerequisites**, **dependencies**, **compilation build order**, or **task scheduling**.
* Identifying whether a cyclic dependency exists among tasks.

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

        // Build graph and calculate in-degrees
        for (int[] edge : prerequisites) {
            int course = edge[0];
            int prereq = edge[1];
            adj.get(prereq).add(course);
            inDegree[course]++;
        }

        // Add all courses with 0 prerequisites to the queue
        Queue<Integer> queue = new ArrayDeque<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }

        int[] order = new int[numCourses];
        int index = 0;

        while (!queue.isEmpty()) {
            int current = queue.poll();
            order[index++] = current;

            // Reduce in-degree for all neighboring courses
            for (int neighbor : adj.get(current)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }

        // If cycle exists, index will be less than total courses
        return (index == numCourses) ? order : new int[0];
    }
}
```

---

## 18.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(V + E)$ where $V$ is vertices/courses and $E$ is edges/prerequisites.
* **Space Complexity:** $\mathcal{O}(V + E)$ to store the adjacency list graph and in-degree array.

---

## 18.5 - Classic LeetCode Problems

* **Course Schedule I** (LeetCode #207)
* **Course Schedule II** (LeetCode #210)
* **Alien Dictionary** (LeetCode #269)
* **Minimum Height Trees** (LeetCode #310)
* **Sequence Reconstruction** (LeetCode #444)
* **Parallel Courses** (LeetCode #1136)

---

## 18.6 - Sources used for this file:
https://en.wikipedia.org/wiki/Topological_sorting <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d56408d2bb2d978e2d42 <br>
https://www.geeksforgeeks.org/topological-sorting/ <br>
https://techinterviewhandbook.org/algorithms/graph/
