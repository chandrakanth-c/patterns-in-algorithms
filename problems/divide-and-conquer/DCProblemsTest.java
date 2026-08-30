import java.util.Arrays;

public class DCProblemsTest {

    public static void main(String[] args) {
        testMaxSubarray();
        testStrassenMultiply2x2();
        testQuickSort();
        testRandomizedSelect();
        testClosestPair();
        System.out.println("ALL DIVIDE AND CONQUER TESTS PASSED!");
    }

    private static void assertTrue(boolean condition, String message) {
        if (!condition) {
            throw new AssertionError("Test Failed: " + message);
        }
    }

    private static void testMaxSubarray() {
        int[] nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        assertTrue(DCProblems.maxSubArray(nums) == 6, "maxSubArray");
    }

    private static void testStrassenMultiply2x2() {
        int[][] A = {{1, 2}, {3, 4}};
        int[][] B = {{5, 6}, {7, 8}};
        int[][] C = DCProblems.strassenMultiply2x2(A, B);
        assertTrue(C[0][0] == 19 && C[0][1] == 22 && C[1][0] == 43 && C[1][1] == 50, "strassen2x2");
    }

    private static void testQuickSort() {
        int[] nums = {9, 7, 5, 11, 12, 2, 14, 3, 10, 6};
        int[] expected = {2, 3, 5, 6, 7, 9, 10, 11, 12, 14};
        DCProblems.quickSort(nums, 0, nums.length - 1);
        assertTrue(Arrays.equals(nums, expected), "quickSort");
    }

    private static void testRandomizedSelect() {
        int[] nums = {9, 7, 5, 11, 12, 2, 14, 3, 10, 6};
        int kth = DCProblems.randomizedSelect(nums, 0, nums.length - 1, 5);
        assertTrue(kth == 7, "randomizedSelect");
    }

    private static void testClosestPair() {
        DCProblems.Point[] pts = {
            new DCProblems.Point(2, 3),
            new DCProblems.Point(12, 30),
            new DCProblems.Point(40, 50),
            new DCProblems.Point(5, 1),
            new DCProblems.Point(12, 10),
            new DCProblems.Point(3, 4)
        };
        double d = DCProblems.closestPair(pts);
        assertTrue(Math.abs(d - Math.sqrt(2)) < 1e-4, "closestPair");
    }
}
