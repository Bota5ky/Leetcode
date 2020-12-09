package leetcodejava;

//https://leetcode-cn.com/problems/guess-number-higher-or-lower/
public class leetcode374 extends GuessGame {
    public int guessNumber(int n) {
        for (int i = 1, j = n; i <= j;) {
            int mid = i + (j - i) / 2;
            if (guess(mid) < 0) {
                j = mid - 1;
            } else if (guess(mid) > 0) {
                i = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;
    }
}

class GuessGame {
    int n;

    int guess(int num) {
        if (num > n) {
            return 1;
        } else if (num < n) {
            return -1;
        } else {
            return 0;
        }
    }
}