#include <stdio.h>
#include <pthread.h>

pthread_mutex_t mutex1;
pthread_mutex_t mutex2;

void* thread1_function(void *arg) {
  pthread_mutex_lock(&mutex1);
  printf("Thread 1 acquired mutex1\n");
  sleep(1); // Sleep for 10 seconds to simulate some work...
  
  printf("Thread 1 trying acquire mutex2\n");
  pthread_mutex_lock(&mutex2);
  printf("Thread 1 acquired mutex2\n");

  pthread_mutex_unlock(&mutex2);
  pthread_mutex_unlock(&mutex1);
}

void* thread2_function(void *arg) {
  pthread_mutex_lock(&mutex2);
  printf("Thread 2 acquired mutex2\n");
  sleep(1); // Sleep for 10 seconds to simulate some work

  printf("Thread 2 trying acquire mutex1\n");
  pthread_mutex_lock(&mutex2);
  printf("Thread 2 acquired mutex1\n");

  pthread_mutex_unlock(&mutex1);
  pthread_mutex_unlock(&mutex2);
}

int main() {
  pthread_t thread1, thread2;

  pthread_create(&thread1, NULL, thread1_function, NULL);
  pthread_create(&thread2, NULL, thread2_function, NULL);

  // Wait for both threads to finish;
  pthread_join(thread1, NULL);
  pthread_join(thread2, NULL);

  return 0;
}

