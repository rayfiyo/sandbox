#include <stdio.h>

void func(void){
	printf("ポインタで呼び出す");
}

int main(){
	void (*p)(void);
	p = &func;
	(*p)();
}
