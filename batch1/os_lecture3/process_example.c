#include <stdio.h>
#include <stdlib.h>

int main() {
  printf("PID: %d, PPID: %d\n", getpid(), getppid());
  return 0;
}
