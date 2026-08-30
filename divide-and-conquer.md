# 2 - Divide and Conquer

## 2.1 - Overview (CLRS Chapter 4)

* The **Divide-and-Conquer** paradigm is a recursive problem-solving strategy that operates in three distinct phases at each level of recursion:
  1. **Divide** the problem into one or more subproblems that are smaller instances of the exact same problem.
  2. **Conquer** the subproblems by solving them recursively. If the subproblem sizes are small enough (the *base case*), solve the subproblems in a straightforward manner.
  3. **Combine** the solutions to the subproblems into the solution for the original problem.

---

## 2.2 - Recurrence Relations and the Master Theorem (CLRS 4.5)

The running time of a divide-and-conquer algorithm is characterized by a recurrence relation. When an algorithm divides a problem of size $n$ into $a$ subproblems, each of size $n/b$, with $\mathcal{O}(f(n))$ cost to divide and combine:

$$T(n) = a T\left(\frac{n}{b}\right) + f(n)$$

Where $a \ge 1$ and $b > 1$ are constants.

### The Master Theorem (CLRS 4th Ed. Theorem 4.1)
Let $c_{\text{crit}} = \log_b a$ be the critical exponent.
1. **Case 1 (Leaves Dominate):** If $f(n) = \mathcal{O}(n^{\log_b a - \epsilon})$ for some constant $\epsilon > 0$, then:
   $$T(n) = \Theta(n^{\log_b a})$$
2. **Case 2 (Even Balance across Levels):** If $f(n) = \Theta(n^{\log_b a} \lg^k n)$ for some $k \ge 0$, then:
   $$T(n) = \Theta(n^{\log_b a} \lg^{k+1} n)$$
3. **Case 3 (Root Dominates):** If $f(n) = \Omega(n^{\log_b a + \epsilon})$ for some constant $\epsilon > 0$, and if $a f(n/b) \le c f(n)$ for some constant $c < 1$ and sufficiently large $n$ (regularity condition), then:
   $$T(n) = \Theta(f(n))$$

---

## 2.3 - Classic Example: Merge Sort

* **Divide:** Compute midpoint $q = \lfloor (p + r)/2 \rfloor$ in $\mathcal{O}(1)$.
* **Conquer:** Recursively solve two subproblems of size $n/2$, costing $2T(n/2)$.
* **Combine:** Merge two sorted subarrays of total size $n$ in $\Theta(n)$.
* **Recurrence:** $T(n) = 2T(n/2) + \Theta(n) \implies a=2, b=2, \log_2 2 = 1 \implies T(n) = \Theta(n \lg n)$ by Case 2.

---

### Java Implementation

```java
public class MergeSort {

    public static void sort(int[] arr, int left, int right) {
        if (left < right) {
            // Avoid potential integer overflow with (left + right) / 2
            int mid = left + (right - left) / 2;

            sort(arr, left, mid);
            sort(arr, mid + 1, right);
            merge(arr, left, mid, right);
        }
    }

    private static void merge(int[] arr, int left, int mid, int right) {
        int n1 = mid - left + 1;
        int n2 = right - mid;

        int[] L = new int[n1];
        int[] R = new int[n2];

        System.arraycopy(arr, left, L, 0, n1);
        System.arraycopy(arr, mid + 1, R, 0, n2);

        int i = 0, j = 0, k = left;
        while (i < n1 && j < n2) {
            if (L[i] <= R[j]) {
                arr[k++] = L[i++];
            } else {
                arr[k++] = R[j++];
            }
        }

        while (i < n1) arr[k++] = L[i++];
        while (j < n2) arr[k++] = R[j++];
    }
}
```

---

### Go Implementation

```go
package main

// MergeSort sorts a slice of integers using the Divide-and-Conquer paradigm
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	// Divide and conquer
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Combine
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
```

---

## 2.4 - Time & Space Complexity Analysis

* **Merge Sort Complexity:**
  * **Best, Average, Worst Time:** $\Theta(n \log n)$ universally guaranteed.
  * **Auxiliary Space:** $\Theta(n)$ for auxiliary merge buffers + $\mathcal{O}(\log n)$ stack frames.
* **Divide & Conquer vs Dynamic Programming:**
  * Divide & conquer solves **disjoint** subproblems and does not cache results.
  * DP solves **overlapping** subproblems by caching and reusing subproblem solutions.

---

## 2.5 - Classic CLRS Algorithms & LeetCode Benchmarks

* **Maximum-Subarray Problem** (CLRS 4.1 / LeetCode #53)
* **Strassen's Matrix Multiplication** (CLRS 4.2) — $\mathcal{O}(n^{\log_2 7}) \approx \mathcal{O}(n^{2.807})$
* **Quicksort (Hoare / Lomuto Partitioning)** (CLRS Chapter 7 / LeetCode #912)
* **K-th Order Statistic Selection (Randomized-Select)** (CLRS Chapter 9 / LeetCode #215)
* **Closest Pair of 2D Points** (CLRS Chapter 33)

---

## 2.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 4: Divide-and-Conquer (pp. 76–122)
  * Section 4.1: The maximum-subarray problem (pp. 77–88)
  * Section 4.2: Strassen's algorithm for matrix multiplication (pp. 89–98)
  * Section 4.5: The master method for solving recurrences (pp. 110–118)
* https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm
* https://www.geeksforgeeks.org/divide-and-conquer-algorithm-introduction/
