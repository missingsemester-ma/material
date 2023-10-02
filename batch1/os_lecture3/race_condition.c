#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define NUM_THREADS 2
#define ITERATIONS 100000 

int shared_variable = 0;

void* thread_function(void* arg) {
  for (int i = 0; i < ITERATIONS; ++i) {
    shared_variable++;
  }

  return NULL;
}

int main() {
  pthread_t threads[NUM_THREADS];

  for (int i = 0; i < NUM_THREADS; ++i) {
    if (pthread_create(&threads[i], NULL, thread_function, NULL) != 0) {
      perror("pthread_create");
      exit(EXIT_FAILURE);
    }
  }

  for (int i = 0; i < NUM_THREADS; ++i) {
    if (pthread_join(threads[i], NULL) != 0) {
      perror("pthread_join");
      exit(EXIT_FAILURE);
    }
  }

  printf("Final value of the shared variable: %d\n", shared_variable);

  return 0;
}

