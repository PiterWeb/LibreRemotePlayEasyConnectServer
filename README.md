# LibreRemotePlay EasyConnect Server

## Description

LibreRemotePlay EasyConnect Server is a server that allows you to connect between devices using LibreRemote Play. It is designed to be easy to set up and self-host, providing a simple way to connect devices without the need to passing codes between devices.

## Features

- Easy to set up and use it, is a binary that can be run on any platform that Go supports.
- Self-host first (Includes a Dockerfile)

## Usage with Docker

1. Build the Docker image:

Make sure you have Docker or Podman installed and running on your machine. Then, navigate to the directory containing the `Dockerfile` and run the following command to build the Docker image:

- Docker:
```bash
docker build -t lrpec:latest .
```

- Podman:
```bash
podman build -t lrpec:latest .
```

This command will create a Docker image named `lrpec` with the `latest` tag.

2. Run the Docker container:

Use the following command to run the container, replacing `<port>` with the desired port number:

- Docker:
```bash
docker run -p <port>:80 --name lrpec lrpec:latest
```

- Podman:
```bash
podman run -p <port>:80 --name lrpec lrpec:latest
```
This command will start the container and map the specified port on your host machine to port 80 in the container.

# Usage without Docker

1. Donwload Go from https://go.dev/dl/

2. Clone the repository:
```bash
git clone https://github.com/PiterWeb/LibreRemotePlayEasyConnectServer
cd LibreRemotePlayEasyConnectServer
```

3. Build the project:
```bash
go build -o lrpec
```
4. Run the server:
```bash
./lrpec <port>
```

Replace `<port>` with the desired port number. The server will start and listen for incoming connections on the specified port.

