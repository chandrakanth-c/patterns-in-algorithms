import java.util.*;

public class BitManipulation {

    // 1. SingleNumber — element appearing once; others appear twice (LeetCode #136)
    public static int singleNumber(int[] nums) {
        int result = 0;
        for (int n : nums) result ^= n;
        return result;
    }

    // 2. SingleNumberII — element appearing once; others appear three times (LeetCode #137)
    // Bit counting modulo 3 using two bitmasks.
    public static int singleNumberII(int[] nums) {
        int ones = 0, twos = 0;
        for (int n : nums) {
            ones = (ones ^ n) & ~twos;
            twos = (twos ^ n) & ~ones;
        }
        return ones;
    }

    // 3. SingleNumberIII — two elements each appearing once; others twice (LeetCode #260)
    public static int[] singleNumberIII(int[] nums) {
        int xor = 0;
        for (int n : nums) xor ^= n;
        int diff = xor & (-xor); // isolate rightmost set bit
        int a = 0;
        for (int n : nums) {
            if ((n & diff) != 0) a ^= n;
        }
        int b = xor ^ a;
        return a < b ? new int[]{a, b} : new int[]{b, a};
    }

    // 4. HammingWeight — count set bits (LeetCode #191)
    public static int hammingWeight(int n) {
        int count = 0;
        while (n != 0) {
            n &= (n - 1); // clear lowest set bit
            count++;
        }
        return count;
    }

    // 5. CountBits — dp[i] = number of 1-bits in i (LeetCode #338)
    public static int[] countBits(int n) {
        int[] dp = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            dp[i] = dp[i >> 1] + (i & 1);
        }
        return dp;
    }

    // 6. SubsetsBitmask — all subsets via bitmask enumeration
    public static List<List<Integer>> subsetsBitmask(int[] nums) {
        int n = nums.length;
        int total = 1 << n;
        List<List<Integer>> result = new ArrayList<>();
        for (int mask = 0; mask < total; mask++) {
            List<Integer> subset = new ArrayList<>();
            for (int i = 0; i < n; i++) {
                if ((mask & (1 << i)) != 0) subset.add(nums[i]);
            }
            result.add(subset);
        }
        return result;
    }
}
