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

### 20.5.1 - Single Number (LeetCode #136)

#### 1. Problem Statement
Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one.

#### 2. Solution Link
* [Go Implementation](problems/bit-manipulation/bit_manipulation_problems.go) (Function: `SingleNumber`)
* [Java Implementation](problems/bit-manipulation/BitManipulation.java) (Method: `singleNumber`)

#### 3. Explanation
The solution utilizes the XOR bitwise operator. The XOR operator has two key properties: $x \oplus x = 0$ and $x \oplus 0 = x$. By XORing all numbers in the array, the elements that appear twice will cancel each other out ($x \oplus x = 0$), leaving only the element that appears once.

#### 4. Conceptual Link to Bit Manipulation
This is the quintessential example of the **Self-inverse / Cancellation property** of XOR. It demonstrates how bitwise operations can solve a problem in $\mathcal{O}(n)$ time and $\mathcal{O}(1)$ space that would otherwise require $\mathcal{O}(n)$ space with a hash map.

### 20.5.2 - Single Number II (LeetCode #137)

#### 1. Problem Statement
Given an integer array `nums` where every element appears three times except for one, which appears exactly once. Find the single element and return it.

#### 2. Solution Link
* [Go Implementation](problems/bit-manipulation/bit_manipulation_problems.go) (Function: `SingleNumberII`)

#### 3. Explanation
This problem is solved by counting the number of set bits at each bit position. Since every number (except one) appears three times, the sum of bits at each position should be a multiple of 3. Any remainder at a position belongs to the single number. The implementation uses two bitmasks (`ones` and `twos`) to track bits that have appeared once or twice.

#### 4. Conceptual Link to Bit Manipulation
It extends the idea of state representation using bits. Instead of a simple boolean state (present or not), we track a state modulo 3 using bitwise logic gates (`AND`, `XOR`, `NOT`).

### 20.5.3 - Single Number III (LeetCode #260)

#### 1. Problem Statement
Given an integer array `nums`, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

#### 2. Solution Link
* [Go Implementation](problems/bit-manipulation/bit_manipulation_problems.go) (Function: `SingleNumberIII`)

#### 3. Explanation
1. XOR all numbers to get `xor = a ^ b`, where `a` and `b` are the unique numbers.
2. Isolate the rightmost set bit in `xor` using `diff = xor & -xor`. This bit must be different between `a` and `b`.
3. Partition the numbers into two groups based on this bit and XOR each group. One group will yield `a` and the other `b`.

#### 4. Conceptual Link to Bit Manipulation
Uses the **Isolate lowest set bit** trick (`x & -x`) to create a partition criteria. It demonstrates how to use a specific bit as a "filter" to separate intertwined signals.

### 20.5.4 - Number of 1 Bits (LeetCode #191)

#### 1. Problem Statement
Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

#### 2. Solution Link
* [Go Implementation](problems/bit-manipulation/bit_manipulation_problems.go) (Function: `HammingWeight`)
* [Java Implementation](problems/bit-manipulation/BitManipulation.java) (Method: `countSetBits`)

#### 3. Explanation
The algorithm uses **Brian Kernighan’s trick**: `n & (n - 1)`. This operation always clears the least significant bit that is set to 1. By repeating this until `n` becomes 0, the number of iterations equals the number of set bits.

#### 4. Conceptual Link to Bit Manipulation
Demonstrates the efficiency of **Brian Kernighan’s algorithm**. Instead of checking all 32 or 64 bits, it only performs operations proportional to the number of set bits $k$, resulting in $\mathcal{O}(k)$ complexity.

---

## 20.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 31: Number-Theoretic Algorithms (Binary representation pp. 902–910)
  * Appendix B: Sets, Relations, and Functions (pp. 1162–1170)
* https://leetcode.com/problems/single-number/
* https://graphics.stanford.edu/~seander/bithacks.html
