#include "stdio.h"

int main(void) {
  printf("hint: %.60lf\n", 0.2);

  for (double i = -2.0; i <= 3.2; i += 0.2) {
    printf("%.60lf\n", i);
  }

  return 0;
}
