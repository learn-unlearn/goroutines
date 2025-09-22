## ScanScope [just a random name to give the project]


### Docker

    ```
    docker run -it -v ./:/usr/src/app -w /usr/src/app golang bash ====== this runs as root
    ```

## Run in container

    ```
    - export GOCACHE=/tmp/go-build // run this in docker container
    - apt-get update && apt-get install -y iproute2
    - apt-get install -y iputils-ping
    - go run main.go
    - go build -o scanner || go build
    ```

### System commands to check your cpu and more

    - nproc                            # CPU cores
    - ulimit -n                        # max open files for the shell
    - cat /proc/sys/fs/file-max        # system-wide FD limit

# While running scan, in another terminal:

    - ss -s                            # socket summary
    - ss -tn state established         # established TCP sockets
    - watch -n1 'ss -s; echo; ls /proc/$(pgrep yourbinary)/fd | wc -l'   # monitor sockets + fds