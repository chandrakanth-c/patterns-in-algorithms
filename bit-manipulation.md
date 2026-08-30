# 20 - Bit Manipulation & Bitmasking

## 20.1 - Overview

* **Bit Manipulation** involves the algorithmic manipulation of bits or other pieces of data shorter than a word.
* Core bitwise operators:
  * `&` (AND), `|` (OR), `^` (XOR), `~` (NOT), `<<` (Left shift), `>>` (Signed right shift), `>>>` (Unsigned right shift).
* Key Bitwise Identities:
  * `x ^ x = 0` and `x ^ 0 = x` (Self-cancellation)
  * `x & (x - 1)` clears the lowest set bit (used in Brian Kernighan’s algorithm).
  * `x & (-x)` extracts the lowest set bit (used in Fenwick Trees).
  * `(x & (1 << i)) != 0` checks if $i$-th bit is set.
  * `x | (1 << i)` sets the $i$-th bit.
  * `x & ~(1 << i)` unsets the $i$-th bit.
  * `x ^ (1 << i)` toggles the $i$-th bit.

---

## 20.2 - Bitmasking in State Representation

* An integer can represent a set of up to 32 (or 64 for `long`) boolean items.
* For example, `mask = 5` in binary is `101_2`, representing that items `0` and `2` are included, while item `1` is excluded.
* Extremely useful in **Bitmask DP** (e.g. Traveling Salesperson Problem, Assignment Problems) where $N \le 20$.

---

## 20.3 - Classic Example: Single Number (XOR Property)

### Java Implementation

```java
public class BitManipulation {

    // 1. Find the element that appears once where all other elements appear twice
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num; // All duplicate pairs cancel out (a ^ a = 0)
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

## 20.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(1)$ for bitwise arithmetic operations; $\mathcal{O}(n)$ to scan an array using bitwise aggregators.
* **Space Complexity:** $\mathcal{O}(1)$ auxiliary space.

---

## 20.5 - Classic LeetCode Problems

* **Single Number I, II & III** (LeetCode #136, #137, #260)
* **Number of 1 Bits** (Hamming Weight) (LeetCode #191)
* **Counting Bits** (LeetCode #338)
* **Reverse Bits** (LeetCode #190)
* **Subsets using Bit Manipulation** (LeetCode #78)
* **Find the Duplicate Number** (LeetCode #287)

---

## 20.6 - Sources used for this file:
https://leetcode.com/problems/single-number/ <br>
https://www.geeksforgeeks.org/bits-manipulation-important-tactics/ <br>
https://graphics.stanford.edu/~seander/bithacks.html <br>
https://techinterviewhandbook.org/algorithms/binary/
