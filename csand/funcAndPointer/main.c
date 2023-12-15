#include <stdio.h>

void func1(void) { printf("天才"); }
void func2(void) { printf("凡人"); }
void func3(void) { printf("ごみ"); }

int main(void) {
  void (*p[3])(void);
  p[0] = &func1;
  p[1] = &func2;
  p[2] = &func3;

  int random = 1; // 0~2のランダム
  (*p[random])();
}
