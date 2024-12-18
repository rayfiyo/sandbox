#include "stdio.h"

void func1(int *n) {
  printf("%d\n", *n);
  *n = 1;
  printf("%d\n", *n);
}

int main(void) {
  int *n;
  printf("%d\n", *n);

  func1(n);

  return 0;
}
