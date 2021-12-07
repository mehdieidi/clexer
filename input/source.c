float time_diff(struct timeval *start, struct timeval *end);

int main(int argc, char **argv) {          
    int pipe_fd[2];

    if (pipe(pipe_fd) == -1) {
        printf("error - pipe failed");
        return errno;
    }

    int pid = fork();

    if (pid == -1) {
        printf("rrror - fork failed");
        return errno;
    } else if (pid == 0) {
        struct timeval cur;
        gettimeofday(&cur, NULL);
        
        write(pipe_fd[1], &cur, sizeof(cur));

        execvp(argv[1], &argv[1]);

        exit(0);
    } else {
        wait(NULL);

        struct timeval cur;
        gettimeofday(&cur, NULL);

        char buf[20];
        read(pipe_fd[0], buf, sizeof(buf));

        printf("\n\nElapsed Time: %f seconds\n", time_diff(buf, &cur));

        close(pipe_fd[0]);
        close(pipe_fd[1]);
    }

    return 0;
}

// time_diff returns the difference of start and end times.
float time_diff(struct timeval *start, struct timeval *end){
    return (end->tv_sec - start->tv_sec) + 1e-6*(end->tv_usec - start->tv_usec);
}