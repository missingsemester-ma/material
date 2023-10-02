#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

#define NUM_THREADS 2
#define ITERATIONS 1000000

// Structure to encapsulate shared data and synchronization
typedef struct {
  int shared_variable;
  pthread_mutex_t mutex;
  pthread_cond_t cond;
} Monitor;

Monitor monitor = {0, PTHREAD_MUTEX_INITIALIZER, PTHREAD_COND_INITIALIZER};

void* thread_function(void* arg) {
  for (int i = 0; i < ITERATIONS; ++i) {
    pthread_mutex_lock(&monitor.mutex);
    monitor.shared_variable++;
    pthread_cond_signal(&monitor.cond);
    pthread_mutex_unlock(&monitor.mutex);
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

  printf("Final value of the shared variable: %d\n", monitor.shared_variable);

  return 0;
}
