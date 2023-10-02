#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

#define NUM_THREADS 2

void* thread_function(void* arg) {
  int thread_id = *(int*) arg;
  printf("Hello from Thread %d\n", thread_id);
  printf("PID: %d, PPID: %d\n", getpid(), getppid());
}

int main() {
  pthread_t threads[NUM_THREADS];
  int thread_args[NUM_THREADS];

  printf("Main thread starting PID:%d, PPID%d\n", getpid(), getppid());

  for (int i = 0; i < NUM_THREADS; ++i) {
    thread_args[i] = i + 1;
    if (pthread_create(
      &threads[i], NULL, thread_function, (void*)&thread_args[i]) != 0) {

      perror("pthread_creater");
      exit(EXIT_FAILURE);
    }
  }

  for (int i = 0 ; i < NUM_THREADS; ++i) {
    if (pthread_join(threads[i], NULL) != 0) {
      perror("pthread_join");
      exit(EXIT_FAILURE);
    }
  }


  printf("Main thread will exit in 10 seconds\n");
  printf("Main thread exiting. \n");
  return 0;
}
