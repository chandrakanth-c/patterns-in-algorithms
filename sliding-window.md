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

### 4.5.1 - Longest Substring Without Repeating Characters (LeetCode #3)

#### 1. Problem Statement
Given a string `s`, find the length of the longest substring without repeating characters.

#### 2. Solution Link
* [Go Implementation](problems/sliding-window/sliding_window_problems.go) (Function: `LengthOfLongestSubstring`)

#### 3. Explanation
The algorithm uses a map to store the last seen index of each character. As we iterate with the `right` pointer, if we encounter a character already in the map and its index is within the current window, we move the `left` pointer to `index + 1`.

#### 4. Conceptual Link to Sliding Window
This problem is a classic application of the dynamic sliding window. The window size grows as long as characters are unique and shrinks (by jumping `left`) when a duplicate is found, maintaining the "no duplicates" invariant.

### 4.5.2 - Minimum Window Substring (LeetCode #76)

#### 1. Problem Statement
Given two strings `s` and `t`, return the minimum window substring of `s` such that every character in `t` (including duplicates) is included in the window. If there is no such substring, return the empty string `""`.

#### 2. Solution Link
* [Go Implementation](problems/sliding-window/sliding_window_problems.go) (Function: `MinWindow`)

#### 3. Explanation
We use two frequency maps: one for the characters required from `t` and another for the current window. We expand the `right` pointer until the window contains all characters from `t`. Then, we contract the `left` pointer as much as possible while still maintaining all characters from `t` to find the minimum length.

#### 4. Conceptual Link to Sliding Window
This illustrates a "catch-up" sliding window where the `right` pointer aggressively finds a valid state, and the `left` pointer tentatively optimizes it, searching for the minimal valid range.

### 4.5.3 - Longest Repeating Character Replacement (LeetCode #424)

#### 1. Problem Statement
You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times. Return the length of the longest substring containing the same letter you can get after performing the above operations.

#### 2. Solution Link
* [Go Implementation](problems/sliding-window/sliding_window_problems.go) (Function: `CharacterReplacement`)

#### 3. Explanation
We maintain a window and track the frequency of characters within it. The key invariant is that the number of replacements needed (window size minus the frequency of the most frequent character) must be $\le k$. If it exceeds $k$, we shrink the window from the left.

#### 4. Conceptual Link to Sliding Window
This problem uses the window to track a global state (max frequency). It demonstrates how sliding window can handle "allowed violations" (up to $k$ replacements) by maintaining a count of the elements that don't match the majority.

### 4.5.4 - Max Consecutive Ones III (LeetCode #1004)

#### 1. Problem Statement
Given a binary array `nums` and an integer `k`, return the maximum number of consecutive `1`'s in the array if you can flip at most `k` `0`'s.

#### 2. Solution Link
* [Go Implementation](problems/sliding-window/sliding_window_problems.go) (Function: `MaxConsecutiveOnes`)

#### 3. Explanation
We expand the `right` pointer and count the number of zeros encountered. If the zero count exceeds `k`, we increment the `left` pointer until we remove a zero from the window, bringing the count back to `k`.

#### 4. Conceptual Link to Sliding Window
This is a variation of the dynamic sliding window where the "validity" of the window is determined by a resource budget ($k$ flips). The window expands to consume the budget and contracts when bankrupt.

### 4.5.5 - Subarray Product Less Than K (LeetCode #713)

#### 1. Problem Statement
Given an array of integers `nums` and an integer `k`, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than `k`.

#### 2. Solution Link
* [Go Implementation](problems/sliding-window/sliding_window_problems.go) (Function: `NumSubarrayProductLessThanK`)

#### 3. Explanation
We maintain the product of elements in the current window. As `right` advances, we multiply by `nums[right]`. If the product becomes $\ge k$, we divide by `nums[left]` and increment `left`. For each valid window `[left, right]`, all subarrays ending at `right` and starting at or after `left` are valid, adding `right - left + 1` to the total count.

#### 4. Conceptual Link to Sliding Window
Instead of finding the "longest" or "shortest" window, this problem uses the window boundaries to count all valid contiguous sub-ranges. It leverages the monotonicity of the product (with positive integers) to efficiently count subarrays.

### Other Benchmarks
* **Sliding Window Maximum** (LeetCode #239)

---

## 4.6 - Sources used for this file:
* **Cormen, T. H., Leiserson, C. E., Rivest, R. L., & Stein, C. (2022).** *Introduction to Algorithms (4th ed.)*. MIT Press.
  * Chapter 17: Amortized Analysis (Aggregate analysis pp. 488–492)
* https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/
* https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d30608d2bb2d978e1b1d
