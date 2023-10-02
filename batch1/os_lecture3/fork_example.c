#include <stdio.h>
#include <unistd.h>

int main() {
  pid_t pid = fork();

  if (pid == -1) {
    perror("fork");
    return 1;
  } else if (pid > 0) {
    printf("Parent process (PID: %d)\n", getpid());
  } else {
    printf("Child process (PID: %d), Parent Process PID: %d\n", getpid(), getppid());
  }

  return 0;
}
