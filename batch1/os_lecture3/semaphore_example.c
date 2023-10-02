#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <semaphore.h>

#define NUM_THREADS 2
#define ITERATIONS 1000000

int shared_variable = 0;  
sem_t semaphore;

void* thread_function(void* arg) {
  for (int i = 0; i < ITERATIONS; ++i) {
    sem_wait(&semaphore);
    shared_variable++;
    sem_post(&semaphore);
  }

  return NULL;
}

int main() {
  pthread_t threads[NUM_THREADS];

  if (sem_init(&semaphore, 0, 1) != 0) {
    perror("sem_init");
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

  sem_destroy(&semaphore);

  return 0;
}

