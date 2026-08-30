public class DPProblemsTest {

    public static void main(String[] args) {
        testCutRod();
        testMatrixChainOrder();
        testLongestCommonSubsequence();
        testOptimalBST();
        testClimbStairs();
        testCoinChange();
        testLengthOfLIS();
        testMinDistance();
        System.out.println("ALL DYNAMIC PROGRAMMING TESTS PASSED!");
    }

    private static void assertTrue(boolean condition, String message) {
        if (!condition) {
            throw new AssertionError("Test Failed: " + message);
        }
    }

    private static void testCutRod() {
        int[] prices = {1, 5, 8, 9, 10, 17, 17, 20, 24, 30};
        assertTrue(DPProblems.cutRod(prices, 1) == 1, "cutRod 1");
        assertTrue(DPProblems.cutRod(prices, 4) == 10, "cutRod 4");
        assertTrue(DPProblems.cutRod(prices, 8) == 22, "cutRod 8");
    }

    private static void testMatrixChainOrder() {
        int[] p = {30, 35, 15, 5, 10, 20, 25};
        assertTrue(DPProblems.matrixChainOrder(p) == 15125, "matrixChainOrder");
    }

    private static void testLongestCommonSubsequence() {
        assertTrue(DPProblems.longestCommonSubsequence("abcde", "ace") == 3, "lcs 1");
        assertTrue(DPProblems.longestCommonSubsequence("abc", "abc") == 3, "lcs 2");
        assertTrue(DPProblems.longestCommonSubsequence("abc", "def") == 0, "lcs 3");
    }

    private static void testOptimalBST() {
        double[] p = {0.15, 0.10, 0.05, 0.10, 0.20};
        double[] q = {0.05, 0.10, 0.05, 0.05, 0.05, 0.10};
        double cost = DPProblems.optimalBST(p, q, 5);
        assertTrue(Math.abs(cost - 2.75) < 1e-4, "optimalBST");
    }

    private static void testClimbStairs() {
        assertTrue(DPProblems.climbStairs(2) == 2, "climbStairs 2");
        assertTrue(DPProblems.climbStairs(3) == 3, "climbStairs 3");
        assertTrue(DPProblems.climbStairs(5) == 8, "climbStairs 5");
    }

    private static void testCoinChange() {
        assertTrue(DPProblems.coinChange(new int[]{1, 2, 5}, 11) == 3, "coinChange 1");
        assertTrue(DPProblems.coinChange(new int[]{2}, 3) == -1, "coinChange 2");
        assertTrue(DPProblems.coinChange(new int[]{1}, 0) == 0, "coinChange 3");
    }

    private static void testLengthOfLIS() {
        assertTrue(DPProblems.lengthOfLIS(new int[]{10, 9, 2, 5, 3, 7, 101, 18}) == 4, "lis 1");
        assertTrue(DPProblems.lengthOfLIS(new int[]{0, 1, 0, 3, 2, 3}) == 4, "lis 2");
        assertTrue(DPProblems.lengthOfLIS(new int[]{7, 7, 7, 7, 7}) == 1, "lis 3");
    }

    private static void testMinDistance() {
        assertTrue(DPProblems.minDistance("horse", "ros") == 3, "editDistance 1");
        assertTrue(DPProblems.minDistance("intention", "execution") == 5, "editDistance 2");
    }
}
