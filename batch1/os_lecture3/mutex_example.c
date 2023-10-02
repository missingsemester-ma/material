#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define NUM_THREADS 2
#define ITERATIONS 1000000

int shared_variable = 0;  // Shared variable that will lead to a race condition
pthread_mutex_t mutex;    // Mutex to protect the shared variable

void* thread_function(void* arg) {
  for (int i = 0; i < ITERATIONS; ++i) {
    pthread_mutex_lock(&mutex);
    shared_variable++;
    pthread_mutex_unlock(&mutex);
  }

  return NULL;
}

int main() {
  pthread_t threads[NUM_THREADS];

  if (pthread_mutex_init(&mutex, NULL) != 0) {
    perror("pthread_mutex_init");
    exit(EXIT_FAILURE);
  }

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

  pthread_mutex_destroy(&mutex);

  return 0;
}
