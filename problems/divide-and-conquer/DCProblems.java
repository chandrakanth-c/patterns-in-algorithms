import java.util.Arrays;

public class DCProblems {

    // 1. Maximum-Subarray Problem (CLRS 4.1 / LeetCode #53)
    public static int maxSubArray(int[] nums) {
        return findMaxSubarray(nums, 0, nums.length - 1);
    }

    private static int findMaxSubarray(int[] nums, int low, int high) {
        if (low == high) return nums[low];
        int mid = low + (high - low) / 2;
        int leftMax = findMaxSubarray(nums, low, mid);
        int rightMax = findMaxSubarray(nums, mid + 1, high);
        int crossMax = findMaxCrossingSubarray(nums, low, mid, high);
        return Math.max(leftMax, Math.max(rightMax, crossMax));
    }

    private static int findMaxCrossingSubarray(int[] nums, int low, int mid, int high) {
        int leftSum = Integer.MIN_VALUE, sum = 0;
        for (int i = mid; i >= low; i--) {
            sum += nums[i];
            leftSum = Math.max(leftSum, sum);
        }
        int rightSum = Integer.MIN_VALUE;
        sum = 0;
        for (int i = mid + 1; i <= high; i++) {
            sum += nums[i];
            rightSum = Math.max(rightSum, sum);
        }
        return leftSum + rightSum;
    }

    // 2. Strassen's 2x2 Matrix Multiplication (CLRS 4.2)
    public static int[][] strassenMultiply2x2(int[][] A, int[][] B) {
        int p1 = A[0][0] * (B[0][1] - B[1][1]);
        int p2 = (A[0][0] + A[0][1]) * B[1][1];
        int p3 = (A[1][0] + A[1][1]) * B[0][0];
        int p4 = A[1][1] * (B[1][0] - B[0][0]);
        int p5 = (A[0][0] + A[1][1]) * (B[0][0] + B[1][1]);
        int p6 = (A[0][1] - A[1][1]) * (B[1][0] + B[1][1]);
        int p7 = (A[0][0] - A[1][0]) * (B[0][0] + B[0][1]);

        int[][] C = new int[2][2];
        C[0][0] = p5 + p4 - p2 + p6;
        C[0][1] = p1 + p2;
        C[1][0] = p3 + p4;
        C[1][1] = p5 + p1 - p3 - p7;
        return C;
    }

    // 3. Quicksort (CLRS Chapter 7 / LeetCode #912)
    public static void quickSort(int[] nums, int p, int r) {
        if (p < r) {
            int q = partition(nums, p, r);
            quickSort(nums, p, q - 1);
            quickSort(nums, q + 1, r);
        }
    }

    private static int partition(int[] nums, int p, int r) {
        int x = nums[r]; // pivot
        int i = p - 1;
        for (int j = p; j < r; j++) {
            if (nums[j] <= x) {
                i++;
                int temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
            }
        }
        int temp = nums[i + 1];
        nums[i + 1] = nums[r];
        nums[r] = temp;
        return i + 1;
    }

    // 4. Randomized-Select in Expected Linear Time (CLRS 9.2)
    public static int randomizedSelect(int[] nums, int p, int r, int k) {
        if (p == r) return nums[p];
        int q = partition(nums, p, r);
        int rank = q - p + 1;
        if (k == rank) return nums[q];
        else if (k < rank) return randomizedSelect(nums, p, q - 1, k);
        else return randomizedSelect(nums, q + 1, r, k - rank);
    }

    // 5. Closest Pair of Points (CLRS Chapter 33)
    public static class Point {
        public double x, y;
        public Point(double x, double y) { this.x = x; this.y = y; }
    }

    public static double closestPair(Point[] points) {
        Point[] pts = points.clone();
        Arrays.sort(pts, (a, b) -> Double.compare(a.x, b.x));
        return closestUtil(pts, 0, pts.length - 1);
    }

    private static double closestUtil(Point[] pts, int left, int right) {
        if (right - left <= 3) {
            double minD = Double.MAX_VALUE;
            for (int i = left; i <= right; i++) {
                for (int j = i + 1; j <= right; j++) {
                    minD = Math.min(minD, dist(pts[i], pts[j]));
                }
            }
            return minD;
        }

        int mid = left + (right - left) / 2;
        double midX = pts[mid].x;
        double d1 = closestUtil(pts, left, mid);
        double d2 = closestUtil(pts, mid + 1, right);
        double d = Math.min(d1, d2);

        java.util.List<Point> strip = new java.util.ArrayList<>();
        for (int i = left; i <= right; i++) {
            if (Math.abs(pts[i].x - midX) < d) {
                strip.add(pts[i]);
            }
        }
        strip.sort((a, b) -> Double.compare(a.y, b.y));

        for (int i = 0; i < strip.size(); i++) {
            for (int j = i + 1; j < strip.size() && (strip.get(j).y - strip.get(i).y) < d; j++) {
                d = Math.min(d, dist(strip.get(i), strip.get(j)));
            }
        }
        return d;
    }

    private static double dist(Point p1, Point p2) {
        return Math.sqrt((p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y) * (p1.y - p2.y));
    }
}
