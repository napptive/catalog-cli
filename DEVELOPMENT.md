# Development Guide

This document describes the development process of Napptive catalog cli.

## Prerequisites

- Read the [CONTRIBUTING guidelines](CONTRIBUTING.md).
- Golang version 1.19+
    - Install `go1.19` from [official site](https://go.dev/dl/). Unpack the binary and place it somewhere, assume it's in the home path `~/go/`, below is an example command, you should choose the right binary according to your system.
        ```bash
        wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
        tar xzf go1.20.2.linux-amd64.tar.gz
        ```
    - Add `~/go/bin` to your `PATH` and `GOROOT` environment variable.
        ```bash
        export GOROOT=~/go/
        export PATH=$PATH:$GOROOT/bin
        ```
- Component [catalog-manager](https://github.com/napptive/catalog-manager) need to be running.

### Build

There is a `Makefile` provided in the root directory with diffrent targets.

- Build the project files for _Local Env_:
    ```bash
    make build
    ```
- Build the project files for _MacOS_:
    ```bash
    make build-darwin
    ```
- Build the project files for _Linux_:
    ```bash
    make build-linux
    ```

This will create a binary file in the `<current_dir>/build/bin/` directory.

**NOTE**
> If you are developing with MacOS/Darwin, you must install [gnu-sed](https://www.gnu.org/software/sed/).
```bash
brew install gnu-sed
```

### Test

- **Run the Unit test**:
    ```bash
    make test
    ```
### K8s
- **Generate the Kubernetes deployment files:**
    ```bash
    make k8s
    ```
   
### Docker
- Prepare the Dockerfile folder with all the extra files
    ```bash
    make docker-prep
    ```
- Build the docker image
    ```bash
    make docker-build
    ```
- Push the image to the selected repository
    ```bash
    make docker-push
    ```
    > **NOTE**: You must set the `TARGET_DOCKER_REGISTRY` environment variable to the repository you want to push the image to. And you must make login before to push the docker image.

### Clean up
- Clean up all the build files
    ```bash
    make clean
    ```

## Code Layout structure
The code layout structure is based on the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
Please refer to the link for more details about the structure.
 