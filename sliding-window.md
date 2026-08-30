# 4 - Sliding Window

## 4.1 - Overview & Theoretical Foundations

* The **Sliding Window** pattern maintains a dynamic or fixed sub-range (window $[L, R]$) over a sequential data structure (array or string).
* As the right boundary $R$ advances, elements are admitted into the window state. When window invariants are violated, the left boundary $L$ advances to expel elements until the state becomes valid.
* **Amortized Analysis (CLRS Chapter 17):**
  * Even though the algorithm uses nested loops (`for right ... { while (!valid) { left++ } }`), each index from $0$ to $n-1$ is visited by $R$ exactly once and by $L$ at most once.
  * Total state additions $= n$, total state removals $\le n$.
  * By aggregate analysis, the total cost across all iterations is $\Theta(n)$, giving an amortized cost of $\mathcal{O}(1)$ per character.

---

## 4.2 - Properties of a problem that suggests Sliding Window

* Input is a linear collection (array, slice, string).
* The problem asks for the **longest**, **shortest**, **maximum**, or **minimum** contiguous subarray or substring that satisfies a constraint.
* The condition exhibits **monotonicity**: expanding the window makes the condition harder/easier to satisfy in a predictable direction.

---

## 4.3 - Classic Example: Longest Substring with at Most K Distinct Characters

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class SlidingWindow {

    public static int lengthOfLongestSubstringKDistinct(String s, int k) {
        if (s == null || s.length() == 0 || k == 0) return 0;

        int left = 0, maxLength = 0;
        Map<Character, Integer> freqMap = new HashMap<>();

        for (int right = 0; right < s.length(); right++) {
            char rightChar = s.charAt(right);
            freqMap.put(rightChar, freqMap.getOrDefault(rightChar, 0) + 1);

            // Shrink window from the left until frequency map size <= k
            while (freqMap.size() > k) {
                char leftChar = s.charAt(left);
                freqMap.put(leftChar, freqMap.get(leftChar) - 1);
                if (freqMap.get(leftChar) == 0) {
                    freqMap.remove(leftChar);
                }
                left++;
            }

            maxLength = Math.max(maxLength, right - left + 1);
        }

        return maxLength;
    }
}
```

---

### Go Implementation

```go
package main

// LengthOfLongestSubstringKDistinct finds the longest substring containing at most k distinct characters
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 || k == 0 {
		return 0
	}

	left := 0
	maxLength := 0
	freqMap := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		rightChar := s[right]
		freqMap[rightChar]++

		// Shrink window until we satisfy at most k distinct characters
		for len(freqMap) > k {
			leftChar := s[left]
			freqMap[leftChar]--
			if freqMap[leftChar] == 0 {
				delete(freqMap, leftChar)
			}
			left++
		}

		currentWindowLen := right - left + 1
		if currentWindowLen > maxLength {
			maxLength = currentWindowLen
		}
	}

	return maxLength
}
```

---

## 4.4 - Time & Space Complexity Analysis

* **Time Complexity:** $\mathcal{O}(n)$ total amortized time, where $n = \text{len}(s)$.
* **Space Complexity:** $\mathcal{O}(k)$ auxiliary space to store frequency tracking entries (at most $k + 1$ keys).

---

## 4.5 - Classic LeetCode Benchmarks

* **Longest Substring Without Repeating Characters** (LeetCode #3)
* **Minimum Window Substring** (LeetCode #76)
* **Longest Repeating Character Replacement** (LeetCode #424)
* **Max Consecutive Ones III** (LeetCode #1004)
* **Subarray Product Less Than K** (LeetCode #713)
* **Sliding Window Maximum** (LeetCode #239)

---

## 4.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 17: Amortized Analysis (Aggregate analysis pp. 488–492)
* https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/
* https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d30608d2bb2d978e1b1d
