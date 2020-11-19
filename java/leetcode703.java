package java;

import java.util.PriorityQueue;

//https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
public class leetcode703 {
    private PriorityQueue<Integer> queue;
    private int limit;

    public leetcode703(int k, int[] nums) {
        limit = k;
        queue = new PriorityQueue<>(k);
        for (int num : nums) {
            add(num);
        }
    }

    public int add(int val) {
        if (queue.size() < limit) {
            queue.add(val);
        } else if (val > queue.peek()) {
            queue.poll();
            queue.add(val);
        }

        return queue.peek();
    }
}
