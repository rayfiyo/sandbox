#include "stdio.h"

// 指数を計算する関数pow
int mypow(int x, int y) {
        int i = 0; // カウンタ変数
        int result = 1; // 結果を格納する変数
        for (i = 0; i < y; i++) {
                result *= x;
        }
        return result;
}

int main(void) {
	printf("%d", mypow(2, 4));

	return 0;
}
