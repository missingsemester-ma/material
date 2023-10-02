#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <sched.h>
#include <unistd.h>
#include <sys/wait.h>

void* thread_function(void* arg) {
  printf("Hello form the new thread! PID = %d, PPID=%d\n", getpid(), getppid());
  while(1);
}

int main() {
  const int STACK_SIZE = 65536; // 64KB
  char* stack = malloc(STACK_SIZE);

  if (stack == NULL) {
    perror("Failed to allocate stack");
    exit(EXIT_FAILURE);
  }

  int value = 42;

  int clone_flags = CLONE_VM | CLONE_SIGHAND | CLONE_FS | CLONE_FILES;
  pid_t pid = clone(thread_function, stack + STACK_SIZE, clone_flags, &value);

  if (pid == -1) {
    perror("clone");
    exit(EXIT_FAILURE);
  }

  if (waitpid(pid, NULL, 0) == -1) {
    perror("waitpid");
    exit(EXIT_FAILURE);
  }

  printf("PID: %d", getpid());
  printf("Main thread exiting.\n");
  while(1);
  return 0;
}
