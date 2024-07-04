#include "stdio.h"

void swap(int *a, int *b) {
  int temp = *a;
  *a = *b;
  *b = temp;

  return;
}

void output(int n, int a[]) {
  int i;
  for (i = 0; i < n; i++) {
    printf("%d ", a[i]);
  }
  printf("\n");

  return;
}

void quick_sort(int left, int right, int n, int a[]) {
  int p, pivot, somewhere, i;
  if (left < right) {
    somewhere = (left + right) / 2;
    pivot = a[somewhere];
    a[somewhere] = a[left];
    p = left;
    for (i = left + 1; i <= right; i++)
      if (a[i] < pivot) {
        p++;
        swap(&a[p], &a[i]);
      }
    a[left] = a[p];
    a[p] = pivot;
    output(n, a);
    quick_sort(left, p - 1, n, a);
    quick_sort(p + 1, right, n, a);
  }
}

int main() {
  int n, i;
  int a[1000];

  printf("n = ");
  scanf("%d", &n);
  for (i = 0; i < n; i++) {
    printf("a[%d] = ", i);
    scanf("%d", &a[i]);
  }

  quick_sort(0, n - 1, n, a);

  return 0;
}
