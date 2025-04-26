#if !defined(__linux__) && !defined(__unix__) && !defined(__APPLE__)
#error "Shared memory isn't supported on you system, use ipc instead"
#endif

#include <fcntl.h>
#include <stddef.h>
#include <sys/mman.h>
#include <sys/stat.h>

void *attach(char *name, size_t size) {
  int fd = shm_open(name, O_RDWR | O_CREAT, PROT_READ | PROT_WRITE);

  if (fd == -1) {
    return NULL;
  }

  return mmap(NULL, size, PROT_READ | PROT_WRITE, MAP_SHARED | MAP_ANONYMOUS,
              fd, 0);
}

void detach(char *name, void *addr, size_t size) {
  munmap(addr, size);

}
