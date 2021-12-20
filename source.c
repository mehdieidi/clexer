// askjdfklasdjfk
/*jsdfkajkdjf*/
#include <stdio.h>
#include <stdlib.h>
#include <sys/mman.h>
#include <sys/time.h>
/*akdsfjd
asdf
asd
*/
#include <sys/wait.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>

float time_diff(struct timeval *start, struct timeval *end);

int main(int argc, char **argv) {        
    const char *shm_name = "shm";
    
    // create shared memory
    int shm_fd = shm_open(shm_name, O_CREAT | O_RDWR, 0666);
    if (shm_fd < 0) {
        perror("error - shm_open failed\n");
        return errno;
    }

    // set the size of the shm
    if (ftruncate(shm_fd, 200 * sizeof(int)) == -1) {
        perror("error - ftruncate failed\n");
        shm_unlink(shm_name);
        return errno;
    }

    int pid = fork();

    if (pid == -1) {
        perror("error - fork failed\n");
        return errno;
    } else if (pid == 0) {
        // memory map the shm to the child process's address space
        struct timeval *ptr = (struct timeval *) mmap(0, 200 * sizeof(int), PROT_READ | PROT_WRITE, MAP_SHARED, shm_fd, 0);

        if (ptr == MAP_FAILED) {
            perror("error - mmap failed");
            return errno;
        }

        struct timeval cur;
        gettimeofday(&cur, NULL);
        
        memcpy(ptr, &cur, sizeof(cur));
        
        execvp(argv[1], &argv[1]);

        exit(0);
    } else {
        wait(NULL);
        
        // memory map the shm to the parent process's address space
        struct timeval *ptr = (struct timeval *) mmap(0, 200 * sizeof(int), PROT_READ | PROT_WRITE, MAP_SHARED, shm_fd, 0);

        struct timeval cur;
        gettimeofday(&cur, NULL);
        
        printf("\n\nElapsed Time: %f seconds\n", time_diff(ptr, &cur));

        shm_unlink(shm_name);
    }

    return 0;
}

// time_diff returns the difference of start and end times.
float time_diff(struct timeval *start, struct timeval *end){
    return (end->tv_sec - start->tv_sec) + 1e-6* (end->tv_usec - start->tv_usec);
}