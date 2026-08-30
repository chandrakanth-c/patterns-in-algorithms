# 20 - Bit Manipulation & Bitmasking

## 20.1 - Overview & Theoretical Foundations (CLRS Chapter 31 & Appendix B)

* Computer architectures represent integers as fixed-width sequences of $w$ bits ($w = 32$ or $64$).
* Bitwise operators execute directly on binary registers in a single processor instruction cycle ($\mathcal{O}(1)$ time).
* **Fundamental Bitwise Identities & Tricks:**
  * `x ^ x = 0` and `x ^ 0 = x` (Self-inverse / Cancellation property).
  * `x & (x - 1)` clears the lowest set bit (Brian Kernighan’s algorithm).
  * `x & (-x)` extracts the lowest set bit as a power of 2 (used in Fenwick Trees).
  * `(x & (1 << i)) != 0` checks if $i$-th bit is set.
  * `x | (1 << i)` sets the $i$-th bit.
  * `x & ~(1 << i)` clears the $i$-th bit.
  * `x ^ (1 << i)` toggles the $i$-th bit.

---

## 20.2 - Bitmasking as Set State Representation

* An integer $M \in [0, 2^N - 1]$ encodes a subset of an $N$-element universe $\{0, 1, \dots, N-1\}$.
* **Set Operations via Bitwise Operators:**
  * $S \cup T \iff S \mid T$ (Union)
  * $S \cap T \iff S \ \& \ T$ (Intersection)
  * $S \setminus T \iff S \ \& \ (\sim T)$ (Set difference)
  * $|S| \iff \text{popcount}(S)$ (Subset cardinality)
* Crucial for **Bitmask Dynamic Programming** (e.g. Traveling Salesperson Problem, $\mathcal{O}(n^2 2^n)$).

---

## 20.3 - Classic Examples: Single Number & Count Set Bits

### Java Implementation

```java
public class BitManipulation {

    // 1. Find the element appearing once when all others appear twice (XOR)
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num;
        }
        return result;
    }

    // 2. Count total number of set bits (Brian Kernighan's Algorithm)
    public static int countSetBits(int n) {
        int count = 0;
        while (n != 0) {
            n = n & (n - 1); // Clears the lowest set bit
            count++;
        }
        return count;
    }
}
```

---

### Go Implementation

```go
package main

import (
	"math/bits"
)

// SingleNumber finds the unique element among pairs
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// CountSetBits counts the number of 1-bits using Brian Kernighan's method
func CountSetBits(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

// FastPopcount uses CPU intrinsic via standard library
func FastPopcount(n uint) int {
	return bits.OnesCount(n)
}
```

---

## 20.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(1)$ for individual bit operations; $\mathcal{O}(k)$ where $k$ is the number of set bits for Kernighan's method.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 20.5 - Classic LeetCode & Benchmark Problems

* **Single Number I, II & III** (LeetCode #136, #137, #260)
* **Number of 1 Bits (Hamming Weight)** (LeetCode #191)
* **Counting Bits** (LeetCode #338)
* **Reverse Bits** (LeetCode #190)
* **Subsets using Bitmask** (LeetCode #78)
* **Traveling Salesperson Problem (Bitmask DP)**

---

## 20.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 31: Number-Theoretic Algorithms (Binary representation pp. 902–910)
  * Appendix B: Sets, Relations, and Functions (pp. 1162–1170)
* https://leetcode.com/problems/single-number/
* https://graphics.stanford.edu/~seander/bithacks.html
