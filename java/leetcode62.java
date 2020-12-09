package java;

//https://leetcode-cn.com/problems/unique-paths/
class leetcode62 {
    public int uniquePaths(int m, int n) {
        int[][] pos = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                int up = 0, left = 0;
                if (i == 0 && j == 0) {
                    pos[i][j] = 1;
                    continue;
                }
                if (i > 0) {
                    up = pos[i - 1][j];
                }
                if (j > 0) {
                    left = pos[i][j - 1];
                }
                pos[i][j] = up + left;
            }
        }
        return pos[m - 1][n - 1];
    }
}