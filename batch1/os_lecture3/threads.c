#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

void *mythread() {
}

int main() {
  pthread_t p1, p2;
  printf("main: begin\n");
  Pthread_create(&p1, NULL, mythread, "A"); 
  Pthread_create(&p2, NULL, mythread, "B");
  // join waits for the threads to finish
  Pthread_join(p1, NULL); 
  Pthread_join(p2, NULL); 
  printf("main: end\n");
  return 0;
}
