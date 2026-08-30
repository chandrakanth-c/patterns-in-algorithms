# 2 - Divide and Conquer

## 2.1 - Overview

* [Recursion](https://en.wikipedia.org/wiki/Recursion_(computer_science)) is the fundamental prerequisite for this topic.
* **Divide and Conquer (D&C)** is an algorithmic paradigm that breaks down a problem into two or more independent sub-problems of the same or related type, solves each sub-problem recursively, and then combines their results to solve the original problem.
* Unlike [Dynamic Programming](https://github.com/chandrakanth-c/patterns-in-algorithms/blob/main/dynamic-programming.md), subproblems in Divide and Conquer are typically **disjoint** (non-overlapping).

---

## 2.2 - Three Core Steps of Divide and Conquer

1. **Divide:** Break the original problem into smaller subproblems.
2. **Conquer:** Solve the subproblems recursively. If a subproblem is small enough (base case), solve it directly.
3. **Combine:** Merge or combine the solutions of the subproblems into the solution for the original problem.

```
                  [Original Problem: P]
                         /      \
                        /        \
              [Subproblem P1]   [Subproblem P2]
                 /     \            /     \
               P1.1    P1.2       P2.1    P2.2
                \       /           \       /
             [Solution S1]        [Solution S2]
                        \        /
                    [Final Solution S]
```

---

## 2.3 - Classic Example: Merge Sort

Merge Sort is the textbook example of Divide and Conquer:
1. **Divide:** Split the array of size $n$ into two halves of size $n/2$.
2. **Conquer:** Recursively sort both halves.
3. **Combine:** Merge the two sorted halves into a single sorted array.

### Java Implementation

```java
public class MergeSort {

    public static void mergeSort(int[] arr, int left, int right) {
        if (left < right) {
            // Find middle point (avoiding integer overflow)
            int mid = left + (right - left) / 2;

            // Step 1 & 2: Divide and conquer left and right halves
            mergeSort(arr, left, mid);
            mergeSort(arr, mid + 1, right);

            // Step 3: Combine solutions by merging
            merge(arr, left, mid, right);
        }
    }

    private static void merge(int[] arr, int left, int mid, int right) {
        int n1 = mid - left + 1;
        int n2 = right - mid;

        int[] L = new int[n1];
        int[] R = new int[n2];

        for (int i = 0; i < n1; ++i) L[i] = arr[left + i];
        for (int j = 0; j < n2; ++j) R[j] = arr[mid + 1 + j];

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

## 2.4 - Time & Space Complexity

* **Time Complexity:**
  * Recurrence relation: $T(n) = 2T(n/2) + \mathcal{O}(n)$
  * By Master Theorem: $\mathcal{O}(n \log n)$ across worst, average, and best cases.
* **Space Complexity:** $\mathcal{O}(n)$ auxiliary space for temporary merge buffers + $\mathcal{O}(\log n)$ call stack space.

---

## 2.5 - Other Classic Divide and Conquer Algorithms

* **Binary Search:** $T(n) = T(n/2) + \mathcal{O}(1) \implies \mathcal{O}(\log n)$
* **Quick Sort:** $T(n) = T(k) + T(n - k - 1) + \mathcal{O}(n) \implies \mathcal{O}(n \log n)$ average
* **Karatsuba Fast Multiplication:** Multiplies two $n$-digit numbers in $\mathcal{O}(n^{\log_2 3}) \approx \mathcal{O}(n^{1.585})$ instead of $\mathcal{O}(n^2)$
* **Strassen's Matrix Multiplication:** Computes matrix product in $\mathcal{O}(n^{2.807})$ instead of $\mathcal{O}(n^3)$
* **Closest Pair of Points:** Finds Euclidean closest pair among $n$ points in $\mathcal{O}(n \log n)$ instead of $\mathcal{O}(n^2)$

---

## 2.6 - Divide and Conquer vs Dynamic Programming

| Characteristic | Divide and Conquer | Dynamic Programming |
| :--- | :--- | :--- |
| **Subproblem Structure** | Disjoint / Independent | Overlapping |
| **Storage / Caching** | Does not cache intermediate subproblem results | Uses Memoization (Top-down) or Tabulation (Bottom-up) |
| **Use Case** | Sorting (Merge/Quick), Searching, Geometric divide | Optimization problems (Min/Max/Count possibilities) |

---

## 2.7 - Sources used for this file:
https://en.wikipedia.org/wiki/Divide-and-conquer_algorithm <br>
https://www.geeksforgeeks.org/divide-and-conquer-algorithm-introduction/ <br>
https://www.geeksforgeeks.org/merge-sort/ <br>
https://mitpress.mit.edu/9780262046305/introduction-to-algorithms/ <br>
https://en.wikipedia.org/wiki/Master_theorem_(analysis_of_algorithms)
