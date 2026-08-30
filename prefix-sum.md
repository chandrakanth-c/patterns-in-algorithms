# 10 - Prefix Sum & Difference Array

## 10.1 - Overview & Theoretical Foundations (CLRS Chapter 2 & 10)

* **Prefix Sum (Cumulative Sum):** Precomputes an array $P$ where $P[i] = \sum_{j=0}^{i-1} A[j]$ with $P[0] = 0$.
  * Any range sum $\sum_{j=L}^R A[j]$ is computed in $\mathcal{O}(1)$ time using the principle of inclusion-exclusion:
    $$\text{Sum}(L, R) = P[R + 1] - P[L]$$
* **2D Prefix Sum:** Precomputes $P[r][c] = \sum_{i=0}^{r-1} \sum_{j=0}^{c-1} A[i][j]$.
  * Any sub-rectangle sum between $(r_1, c_1)$ and $(r_2, c_2)$ evaluates in $\mathcal{O}(1)$:
    $$\text{RegionSum} = P[r_2+1][c_2+1] - P[r_1][c_2+1] - P[r_2+1][c_1] + P[r_1][c_1]$$
* **Difference Array:** The discrete derivative of an array $D[i] = A[i] - A[i-1]$.
  * Applying a range update $+V$ on $[L, R]$ requires only two point updates: $D[L] \mathrel{+}= V$ and $D[R + 1] \mathrel{-}= V$ in $\mathcal{O}(1)$. Recomputing prefix sums recovers the final array in $\mathcal{O}(n)$.

---

## 10.2 - Properties of a problem that suggests Prefix Sum

* Frequent range sum queries on static 1D or 2D arrays.
* Finding subarrays whose sum equals $K$ using a running prefix sum combined with a Hash Map.
* Performing multiple offline range additions before reading the final state.

---

## 10.3 - Classic Example: Subarray Sum Equals K

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class PrefixSum {

    public static int subarraySum(int[] nums, int k) {
        int count = 0;
        int currentPrefixSum = 0;

        // Map stores: (prefixSum, frequencyCount)
        Map<Integer, Integer> prefixMap = new HashMap<>();
        prefixMap.put(0, 1); // Base case: prefix sum 0 occurs once before index 0

        for (int num : nums) {
            currentPrefixSum += num;

            // If (currentPrefixSum - k) exists, there is a subarray summing to k
            if (prefixMap.containsKey(currentPrefixSum - k)) {
                count += prefixMap.get(currentPrefixSum - k);
            }

            prefixMap.put(currentPrefixSum, prefixMap.getOrDefault(currentPrefixSum, 0) + 1);
        }

        return count;
    }
}
```

---

### Go Implementation

```go
package main

// SubarraySum returns the total number of continuous subarrays whose sum equals k
func SubarraySum(nums []int, k int) int {
	count := 0
	currentPrefixSum := 0

	prefixMap := make(map[int]int)
	prefixMap[0] = 1 // Base case

	for _, num := range nums {
		currentPrefixSum += num

		if freq, exists := prefixMap[currentPrefixSum-k]; exists {
			count += freq
		}

		prefixMap[currentPrefixSum]++
	}

	return count
}
```

---

## 10.4 - Time & Space Complexity Analysis

* **Time Complexity:**
  * Precomputation: $\mathcal{O}(n)$ (1D) or $\mathcal{O}(M \times N)$ (2D).
  * Range Query: $\mathcal{O}(1)$ constant time.
  * Subarray Sum Equals $K$: $\mathcal{O}(n)$ single linear pass with map lookups.
* **Space Complexity:** $\mathcal{O}(n)$ auxiliary space to store prefix sums / hash map.

---

## 10.5 - Classic LeetCode Benchmarks

* **Range Sum Query - Immutable** (LeetCode #303)
* **Range Sum Query 2D - Immutable** (LeetCode #304)
* **Subarray Sum Equals K** (LeetCode #560)
* **Continuous Subarray Sum** (LeetCode #523)
* **Car Pooling** (Difference Array) (LeetCode #1094)
* **Corporate Flight Bookings** (LeetCode #1109)

---

## 10.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 2: Getting Started (Prefix invariants)
* https://leetcode.com/problems/subarray-sum-equals-k/
* https://usaco.guide/silver/prefix-sums
