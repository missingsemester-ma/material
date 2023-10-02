#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

// Structure representing a read-write lock
typedef struct {
  pthread_mutex_t mutex;
  pthread_cond_t readers_cv;
  pthread_cond_t writers_cv;
  int readers_count;
  int writers_count;
  int writers_waiting;
} rw_lock_t;

// Initialize a read-write lock
void rw_lock_init(rw_lock_t *lock) {
  pthread_mutex_init(&lock->mutex, NULL);
  pthread_cond_init(&lock->readers_cv, NULL);
  pthread_cond_init(&lock->writers_cv, NULL);
  lock->readers_count = 0;
  lock->writers_count = 0;
  lock->writers_waiting = 0;
}

// Acquire read lock
void rw_lock_acquire_read(rw_lock_t *lock) {
  pthread_mutex_lock(&lock->mutex);
  
  while (lock->writers_count > 0 || lock->writers_waiting > 0) {
    pthread_cond_wait(&lock->readers_cv, &lock->mutex);
  }

  lock->readers_count++;
  pthread_mutex_unlock(&lock->mutex);
}


// Release read lock
void rw_lock_release_read(rw_lock_t *lock) {
    pthread_mutex_lock(&lock->mutex);

    lock->readers_count--;

    if (lock->readers_count == 0) {
      pthread_cond_signal(&lock->writers_cv);
    }

    pthread_mutex_unlock(&lock->mutex);
}

// Acquire write lock
void rw_lock_acquire_write(rw_lock_t *lock) {
  pthread_mutex_lock(&lock->mutex);

  lock->writers_waiting++;

  while (lock->readers_count > 0 || lock->writers_count > 0) {
    pthread_cond_wait(&lock->writers_cv, &lock->mutex);
  }

  lock->writers_waiting--;
  lock->writers_count++;

  pthread_mutex_unlock(&lock->mutex);
}

// Release write lock
void rw_lock_release_write(rw_lock_t *lock) {
  pthread_mutex_lock(&lock->mutex);

  lock->writers_count--;

  if (lock->writers_waiting > 0) {
    pthread_cond_signal(&lock->writers_cv);
  } else {
    pthread_cond_broadcast(&lock->readers_cv);
  }

  pthread_mutex_unlock(&lock->mutex);
}

void* reader_thread(void* arg) {
  int id = *((int*)arg);
  rw_lock_acquire_read(&my_rwlock);
  printf("Reader %d is reading.\n", id);
  rw_lock_release_read(&my_rwlock);
}

void* writer_thread(void* arg) {
  int id = *((int*)arg);
  rw_lock_acquire_write(&my_rwlock);
  printf("Writer %d is writing.\n", id);
  rw_lock_release_write(&my_rwlock);
}

int main() {
  rw_lock_init(&my_rwlock);

  pthread_t readers[3], writers[2];
  int reader_ids[3] = {1, 2, 3};
  int writer_ids[2] = {1, 2};

  for (int i = 0; i < 3; ++i) {
    pthread_create(&readers[i], NULL, reader_thread, &reader_ids[i]);
  }

  for (int i = 0; i < 2; ++i) {
    pthread_create(&writers[i], NULL, writer_thread, &writer_ids[i]);
  }

  for (int i = 0; i < 3; ++i) {
    pthread_join(readers[i], NULL);
  }

  for (int i = 0; i < 2; ++i) {
    pthread_join(writers[i], NULL);
  }

  return 0;
}

