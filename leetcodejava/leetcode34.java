package leetcodejava;

//https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
public class leetcode34 {
    public int[] searchRange(int[] nums, int target) {
        int[] res = { -1, -1 };
        int i = 0, j = nums.length - 1;
        while (i < j) {
            int mid = (i + j) / 2;
            if (nums[mid] >= target) {
                j = mid;
            } else {
                i = mid + 1;
            }
        }
        int left = i;
        if (j < 0 || nums[left] != target) {
            return res;
        }
        i = 0;
        j = nums.length - 1;
        while (i < j) {
            int mid = (i + j) / 2;
            if (nums[mid] > target) {
                j = mid;
            } else {
                i = mid + 1;
            }
        }
        int right = i;
        if (nums[right] == target) {
            right++;
        }
        res[0] = left;
        res[1] = right - 1;
        return res;
    }
}
