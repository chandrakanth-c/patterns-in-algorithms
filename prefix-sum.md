# 10 - Prefix Sum & Difference Array

## 10.1 - Overview

* **Prefix Sum:** A technique where we precalculate cumulative sums of an array $P[i] = \sum_{j=0}^{i-1} A[j]$ so that any range sum query $\sum_{j=L}^R A[j]$ evaluates in $\mathcal{O}(1)$ time as:
  $$\text{sum}(L, R) = P[R + 1] - P[L]$$
* **Difference Array:** The inverse concept used for batch range updates. If we need to add $V$ to all elements in range $[L, R]$, we record $D[L] \mathrel{+}= V$ and $D[R + 1] \mathrel{-}= V$ in $\mathcal{O}(1)$. A single prefix sum pass recovers the updated array.

---

## 10.2 - Properties of a problem that suggests Prefix Sum

* Frequent range sum queries on an array or 2D matrix.
* Finding subarrays with sum equal to $K$ (using running prefix sum + Hash Map).
* Range addition operations $[L, R, +V]$ executed repeatedly before querying final values.

---

## 10.3 - Classic Example: Subarray Sum Equals K

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class PrefixSum {

    public static int subarraySum(int[] nums, int k) {
        int count = 0;
        int currentSum = 0;

        // Map stores: (prefixSum, occurrenceCount)
        Map<Integer, Integer> prefixSumMap = new HashMap<>();
        prefixSumMap.put(0, 1); // Base case: prefix sum 0 occurs once before index 0

        for (int num : nums) {
            currentSum += num;

            // If (currentSum - k) exists in map, then a subarray summing to k exists
            if (prefixSumMap.containsKey(currentSum - k)) {
                count += prefixSumMap.get(currentSum - k);
            }

            prefixSumMap.put(currentSum, prefixSumMap.getOrDefault(currentSum, 0) + 1);
        }

        return count;
    }
}
```

---

## 10.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n)$ because we iterate through the array once and perform $\mathcal{O}(1)$ map lookups/inserts.
* **Space Complexity:** $\mathcal{O}(n)$ to store up to $n$ unique prefix sums in the map.

---

## 10.5 - Classic LeetCode Problems

* **Range Sum Query - Immutable** (LeetCode #303)
* **Range Sum Query 2D - Immutable** (LeetCode #304)
* **Subarray Sum Equals K** (LeetCode #560)
* **Continuous Subarray Sum** (LeetCode #523)
* **Corporate Flight Bookings** (Difference Array) (LeetCode #1109)
* **Car Pooling** (Difference Array) (LeetCode #1094)

---

## 10.6 - Sources used for this file:
https://leetcode.com/problems/subarray-sum-equals-k/ <br>
https://www.geeksforgeeks.org/prefix-sum-array-implementation-applications-competitive-programming/ <br>
https://www.geeksforgeeks.org/difference-array-range-update-query-o1/ <br>
https://usaco.guide/silver/prefix-sums
