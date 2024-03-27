#include <stdio.h>
int main(void) {
  int *ptr;
  *ptr = 10;

  printf("Value: %d\n", *ptr);

  return 0;
}
