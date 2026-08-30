# 4 - Sliding Window

## 4.1 - Overview

* The **Sliding Window** pattern is used to perform required operations on a specific window size of a given array or linked list, such as finding the longest subarray containing all 1s.
* Windows start from the 1st element and keep shifting right by one element and adjusting its length according to the problem constraints.
* Two main types:
  1. **Fixed Size Window:** The window size $K$ remains constant.
  2. **Dynamic / Variable Size Window:** The window grows (expanding `right`) to include candidates and shrinks (contracting `left`) to maintain valid constraints.

---

## 4.2 - Properties of a problem that suggests Sliding Window

* The problem input is a linear data structure like an **Array**, **String**, or **Linked List**.
* You are asked to find the **longest/shortest substring, subarray, or a target value** within a contiguous segment.
* Brute force approaches require nested loops $\mathcal{O}(n^2)$ or $\mathcal{O}(n \cdot k)$ to evaluate every subsegment.

---

## 4.3 - Classic Example: Longest Substring with at most K Distinct Characters

### Java Implementation

```java
import java.util.HashMap;
import java.util.Map;

public class SlidingWindow {

    public static int findLength(String str, int k) {
        if (str == null || str.length() == 0 || k == 0) return 0;

        int windowStart = 0;
        int maxLength = 0;
        Map<Character, Integer> charFrequencyMap = new HashMap<>();

        for (int windowEnd = 0; windowEnd < str.length(); windowEnd++) {
            char rightChar = str.charAt(windowEnd);
            charFrequencyMap.put(rightChar, charFrequencyMap.getOrDefault(rightChar, 0) + 1);

            // Shrink the sliding window until we are left with 'k' distinct characters
            while (charFrequencyMap.size() > k) {
                char leftChar = str.charAt(windowStart);
                charFrequencyMap.put(leftChar, charFrequencyMap.get(leftChar) - 1);
                if (charFrequencyMap.get(leftChar) == 0) {
                    charFrequencyMap.remove(leftChar);
                }
                windowStart++; // Shrink window from the left
            }

            maxLength = Math.max(maxLength, windowEnd - windowStart + 1);
        }

        return maxLength;
    }
}
```

---

## 4.4 - Time & Space Complexity

* **Time Complexity:** $\mathcal{O}(n + n) = \mathcal{O}(n)$ because each character is processed at most twice (once by `windowEnd` and once by `windowStart`).
* **Space Complexity:** $\mathcal{O}(k)$ auxiliary space for the frequency map storing at most $k + 1$ distinct characters.

---

## 4.5 - Classic LeetCode Problems

* **Maximum Sum Subarray of Size K**
* **Longest Substring Without Repeating Characters** (LeetCode #3)
* **Minimum Window Substring** (LeetCode #76)
* **Longest Repeating Character Replacement** (LeetCode #424)
* **Fruit Into Baskets** (LeetCode #904)
* **Sliding Window Maximum** (LeetCode #239)

---

## 4.6 - Sources used for this file:
https://leetcode.com/explore/interview/card/cheatsheets/720/resources/4723/ <br>
https://www.designgurus.io/course-play/grokking-the-coding-interview/doc/6385d30608d2bb2d978e1b1d <br>
https://www.geeksforgeeks.org/window-sliding-technique/ <br>
https://techinterviewhandbook.org/algorithms/string/
