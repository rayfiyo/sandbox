#include "stdio.h"

int *func1(int *one, int *two) {
  printf("func1: %d %d\n", *one, *two);
  return one;
}

int *func2(int *one, int *two) {
  printf("func2: %d %d\n", *one, *two);
  return two;
}

void noname(int *i1, int *i2, int *one) {
  int two = 2;
  i1 = func1(one, &two); // <-- (1)
  printf("(2)\n");
  i2 = func2(one, &two); // <-- (2)
  printf("(2)\n");
}

int main(void) {
  int *i1;
  int *i2;
  int one = 1;
  {
    noname(i1, i2, &one);
  } // <-- (3)
  printf("(3)\n");
  printf("(4): %d\n", *i1); // <-- (4)
  printf("(5): %d\n", *i2); // <-- (5)
  return 0;
}
