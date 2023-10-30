#include "stdio.h"
int main(void) {
	int hoge[10];
	int i = 0;
	for (i = 0; i < 10; i++) {
		scanf("%d", &hoge[i]);
		printf("%d\n", hoge[i]);
	}
}
